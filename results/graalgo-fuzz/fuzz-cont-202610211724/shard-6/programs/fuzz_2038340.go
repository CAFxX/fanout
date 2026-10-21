package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 int, p1 int, p2 int8) string {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            p1 = fi5
            v6 := []int{(p0 + (fi4 >> ((2) & 63))), 1, (p0 | p1)}
            var v7 int16 = 32767
            if false {
                v8 := GenS1{F0: 242}
                _ = v8
            } else {
                v9 := GenS1{F0: 2}
                var v10 bool = false
                _ = v10
                _ = v9
            }
            switch fi4 {
            case 127:
                var v11 int16 = v7
                _ = v11
            case 32767:
                switch 9223372036854775807 {
                case 32768:
                    if len(v6) > 0 {
                        v6[0] = fi5
                    }
                default:
                    var v12 uint64 = 0
                    _ = v12
                }
            case 42:
                fmt.Println(p1, p1)
            default:
                v13 := GenS1{F0: 769}
                _ = v13
            }
            v14 := GenS2{F0: (true && (0.0 > (0.1 + 1e10))), F1: "a"}
            p2 = (p2 | -128)
            _ = fi5
            _ = v14
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    v15 := []int{-1, 1, 0, (^(p0))}
    var v16 int = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return ("foo bar" + "\n\t")
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 42
    var v3 float32 = 1e10
    v3 += (1.5 / 1.5)
    v17 := []int{365}
    v18 := GenS1{F0: 255}
    v17 = append(v17, -9223372036854775808)
    var v19 float32 = 0.0
    var v20 int = 814
    v21 := GenS2{F0: true, F1: ""}
    if false {
        delete(v1, ("" + "\n\t"))
        for fi22 := 0; fi22 < 9; fi22++ {
            fmt.Println(v3, v2)
            v3 = float32((float64((1.5 - (0.0 - 1.5))) * 1e10))
            v19 += 1.5
            _ = fi22
        }
        var v23 float32 = (-2.25 / v19)
        _ = v23
    } else {
        if (!(1 == 42)) {
            var v24 uint64 = 871
            delete(v1, ("a" + ("\n\t" + (("foo bar" + "hello") + "hello"))))
            _ = v24
        } else {
            v2 = v2
            v3 = 1.5
            var v25 int64 = 42
            _ = v25
        }
        for fi26 := 0; fi26 < 20; fi26++ {
            v3 = v3
            v3 = (v19 + ((((float32(3.14159) / ((1.5 + v19) - 1e10)) / -2.25) - v19) * float32((float64((v19 - (v19 + v3))) - 0.1))))
            v27 := GenS2{F0: true, F1: "a"}
            _ = fi26
            _ = v27
        }
    }
    var v28 int32 = 175
    v2 -= (^(v2))
    v19 = (-2.25 * v19)
    switch -2147483648 {
    case -858:
        var v29 bool = (!false)
        _ = v29
    }
    switch ("a" + ("hello" + "\n\t")) {
    case "foo bar":
        if false {
            fmt.Println(v20, v20, v20)
            v28 += 1
        } else {
            var v30 int32 = 2147483647
            var v31 uint64 = 757
            v32 := GenS1{F0: 391}
            _ = v30
            _ = v31
            _ = v32
        }
    }
    v20 += (v20 + (v20 * (v20 ^ (v20 ^ (v20 ^ 637)))))
    var v33 uint16 = 574
    for fi34 := 0; fi34 < 8; fi34++ {
        if (633 != 297) {
            v17 = append(v17, 32767)
        }
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v33
}
