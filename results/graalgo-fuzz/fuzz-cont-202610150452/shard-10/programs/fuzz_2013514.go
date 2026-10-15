package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string) uint64 {
    v6 := []int{4294967295, -9223372036854775808, -243, 203}
    p0 = (p0 + ((((p0 + p0) + (p0 + "")) + ("a" + "foo bar")) + "a"))
    p0 = p0
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 uint32 = 147
    v2 := []int{2147483647, 577}
    v3 := []string{}
    var v4 uint32 = (^(v1))
    var v5 uint32 = v1
    v7 := GenS1{F0: -55}
    v8 := GenS1{F0: 0}
    var v9 uint = 2
    switch -527 {
    case 446:
        v3 = append(v3, ("foo bar" + ""))
    case -2147483648:
        v1 = (v5 & (v4 >> (((v9 * (v9 * (v9 & 2)))) & 63)))
    case -724:
        if true {
            v10 := GenS1{F0: 42}
            _ = v10
        }
    }
    if (1.5 > 0.0) {
        if false {
            for fi11 := 0; fi11 < 5; fi11++ {
                if (1 < 2) {
                    fmt.Println(v4, fi11, fi11)
                } else {
                    if (42 > 42) {
                        _ = len(v3)
                        v12 := GenS1{F0: 1}
                        var v13 int32 = -1
                        _ = v12
                        _ = v13
                    } else {
                        var v14 int64 = 42
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 uint64 = 0
                        var v17 int64 = 9223372036854775807
                        var v18 string = "foo bar"
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                }
                for fi19 := 0; fi19 < 1; fi19++ {
                    v5 = (v4 - (v1 | (v5 & (v4 / (((v1 ^ (v5 - (v4 % ((v4) | 1))))) | 1)))))
                    _ = fi19
                }
                var v20 int16 = -254
                _ = fi11
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int16 = -539
        _ = fi21
        _ = v22
    }
    var v23 uint32 = (v4 + 2147483648)
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 bool = false
        var v26 uint8 = uint8(v9)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
