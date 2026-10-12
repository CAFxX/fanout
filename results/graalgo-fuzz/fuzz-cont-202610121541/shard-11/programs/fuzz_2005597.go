package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []int{32767, 0, 9223372036854775807}
    v2 := []string{"", ("a" + ("hello" + "foo bar")), "\n\t"}
    var v3 int64 = 9223372036854775807
    switch 32768 {
    case 0:
        v1 = append(v1, 4294967295)
    case -2147483648:
        if (284 > 267) {
            switch "\n\t" {
            case "a":
                switch -131 {
                case -702:
                    v4 := GenS1{F0: -636, F1: 100}
                    _ = v4
                case 0:
                    v3 = (v3 ^ -9223372036854775808)
                default:
                    _ = len(v2)
                }
            case "hello":
                v1 = append(v1, 1)
            }
        } else {
            v5 := GenS1{F0: 42, F1: -49}
            _ = v5
        }
    case -9223372036854775808:
        _ = len(v1)
    }
    v6 := GenS2{F0: 1}
    var v7 int64 = (^(v3))
    switch "\n\t" {
    case "foo bar":
        for fi8 := 0; fi8 < 14; fi8++ {
            v3 = v7
            _ = fi8
        }
    case "hello":
        var v9 uint32 = 65535
        _ = v9
    case "":
        v10 := GenS3{F0: ("hello" + "a"), F1: (v7 / ((v3) | 1))}
        _ = v10
    default:
        v11 := GenS3{F0: (("\n\t" + (("foo bar" + (("hello" + (("hello" + "\n\t") + "")) + "hello")) + (("hello" + "") + ("\n\t" + "hello")))) + "hello"), F1: 42}
        _ = v11
    }
    v7 = (v7 + (v3 * v7))
    v2 = append(v2, ("a" + "a"))
    v12 := GenS3{F0: "\n\t", F1: v3}
    for fi13 := 0; fi13 < 11; fi13++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    v7 = (v3 | 42)
    v14 := GenS4{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
