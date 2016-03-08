// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node SelectStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("SelectStmt")

	if node.All {
		ctx.WriteString("all")
		ctx.WriteString(strconv.FormatBool(node.All))
	}

	if len(node.DistinctClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.DistinctClause.Fingerprint(&subCtx, "DistinctClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("distinctClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.FromClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.FromClause.Fingerprint(&subCtx, "FromClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("fromClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.GroupClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.GroupClause.Fingerprint(&subCtx, "GroupClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("groupClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.HavingClause != nil {
		subCtx := FingerprintSubContext{}
		node.HavingClause.Fingerprint(&subCtx, "HavingClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("havingClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.IntoClause != nil {
		subCtx := FingerprintSubContext{}
		node.IntoClause.Fingerprint(&subCtx, "IntoClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("intoClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.Larg != nil {
		subCtx := FingerprintSubContext{}
		node.Larg.Fingerprint(&subCtx, "Larg")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("larg")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.LimitCount != nil {
		subCtx := FingerprintSubContext{}
		node.LimitCount.Fingerprint(&subCtx, "LimitCount")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("limitCount")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.LimitOffset != nil {
		subCtx := FingerprintSubContext{}
		node.LimitOffset.Fingerprint(&subCtx, "LimitOffset")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("limitOffset")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.LockingClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.LockingClause.Fingerprint(&subCtx, "LockingClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("lockingClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if int(node.Op) != 0 {
		ctx.WriteString("op")
		ctx.WriteString(strconv.Itoa(int(node.Op)))
	}

	if node.Rarg != nil {
		subCtx := FingerprintSubContext{}
		node.Rarg.Fingerprint(&subCtx, "Rarg")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("rarg")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.SortClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.SortClause.Fingerprint(&subCtx, "SortClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("sortClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.TargetList.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.TargetList.Fingerprint(&subCtx, "TargetList")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("targetList")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.ValuesLists) > 0 {
		subCtx := FingerprintSubContext{}
		for _, nodeList := range node.ValuesLists {
			for _, subNode := range nodeList {
				subNode.Fingerprint(&subCtx, "ValuesLists")
			}
		}

		if len(subCtx.parts) > 0 {
			ctx.WriteString("valuesLists")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.WhereClause != nil {
		subCtx := FingerprintSubContext{}
		node.WhereClause.Fingerprint(&subCtx, "WhereClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("whereClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if len(node.WindowClause.Items) > 0 {
		subCtx := FingerprintSubContext{}
		node.WindowClause.Fingerprint(&subCtx, "WindowClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("windowClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}

	if node.WithClause != nil {
		subCtx := FingerprintSubContext{}
		node.WithClause.Fingerprint(&subCtx, "WithClause")

		if len(subCtx.parts) > 0 {
			ctx.WriteString("withClause")
			for _, part := range subCtx.parts {
				ctx.WriteString(part)
			}
		}
	}
}