package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []string{"a", (((((("foo bar" + "a") + "foo bar") + "hello") + "foo bar") + ("a" + (("a" + "a") + ("a" + "foo bar")))) + ("hello" + ""))}
    v2 := []int{-658, 130}
    v3 := []string{"hello"}
    var v4 bool = (0 != 42)
    fmt.Println(v4)
    v5 := GenS1{F0: -1}
    var v6 int = 0
    v7 := GenS1{F0: -1}
    var v8 int32 = -179
    fmt.Println(v8, v4, v8)
    v8 = v8
    var v9 uint8 = 1
    var v10 uint8 = (v9 / ((v9) | 1))
    if len(v2) > 0 {
        v2[0] = v6
    }
    var v11 uint = 42
    for fi12 := 0; fi12 < 1; fi12++ {
        switch v6 {
        case 628:
            switch "hello" {
            case "foo bar":
                var v13 uint8 = (v9 - (v10 & v10))
                _ = v13
            default:
                v6 = fi12
            }
        case 32768:
            fmt.Println(v4, v11, v4)
        default:
            switch v6 {
            case 1:
                v9 = v9
            case 0:
                switch "\n\t" {
                case "a":
                    v8 -= (v8 & (v8 & (v8 & v8)))
                case "foo bar":
                    var v14 int16 = -175
                    _ = v14
                case "hello":
                    v6 = -588
                default:
                    switch ("a" + ("\n\t" + "")) {
                    case "foo bar":
                        for fi15 := 0; fi15 < 8; fi15++ {
                            fmt.Println(v11, v11)
                            _ = fi15
                        }
                    default:
                        var v16 int32 = (v8 - -555)
                        _ = v16
                    }
                }
            default:
                v4 = false
            }
        }
        if v4 {
            var v17 string = (("" + ("foo bar" + "a")) + "\n\t")
            v6 -= v6
            _ = v17
        } else {
            var v18 string = "hello"
            _ = v18
        }
        var v19 uint8 = (v9 % ((v10) | 1))
        _ = fi12
        _ = v19
    }
    v1 = append(v1, "a")
    var v20 bool = (("a" + ("" + "hello")) >= "hello")
    v6 += (v6 + int(v11))
    for fi21 := 0; fi21 < 4; fi21++ {
        v6 += v6
        _ = fi21
    }
    var v22 uint32 = 4294967295
    v23 := GenS1{F0: 490}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
