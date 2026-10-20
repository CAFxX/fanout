package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0() string {
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 int8 = 42
        if true {
            var v7 int8 = (v6 & (^(v6)))
            _ = v7
        } else {
            var v8 uint32 = 910
            if ((((true || true) && (!(false || (!true)))) && true) || ((-2147483648 == fi5) || true)) {
                v9 := []int{fi5}
                switch -2147483648 {
                case -941:
                    if ((v6 & (v6 ^ v6)) == (v6 & (v6 % (((v6 << ((275) & 63))) | 1)))) {
                        v10 := GenS1{F0: (((true && (("hello" == "") || false)) && (42 <= 2147483647)) || (!false)), F1: 1, F2: 4294967295}
                        _ = v10
                    }
                }
                _ = v9
            }
            _ = v8
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 string = "a"
            _ = fi11
            _ = v12
        }
        _ = fi5
        _ = v6
    }
    v13 := GenS2{F0: 32768, F1: 4294967295}
    var v14 uint64 = 18446744073709551615
    _ = v13
    _ = v14
    return "a"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 17
    v3 := []string{"\n\t"}
    var v4 uint16 = 317
    delete(v1, 705)
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            if len(v3) > 0 {
                v3[0] = ""
            }
            v4 = v4
            fmt.Println(fi15)
            _ = fi16
        }
        _ = fi15
    }
    var v17 uint = 135
    v17 = v17
    var v18 string = "a"
    _ = len(v1)
    if len(v3) > 0 {
        v3[0] = v18
    }
    switch 1 {
    case 42:
        var v19 float64 = 1e10
        _ = v19
    case -1:
        if false {
            fmt.Println(v2, v2)
            v18 = ((v18 + (v18 + (v18 + v18))) + "a")
            if (2 > 1) {
                fmt.Println(v18)
                v17 -= v17
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
