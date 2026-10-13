package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int8, p1 int16) int {
    v7 := []int{-164, 1}
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint64 = 0
        if false {
            v9 -= v9
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        if false {
            v7 = append(v7, (^(fi8)))
        } else {
            var v10 int64 = 42
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 32767
}
func genFn1() uint64 {
    v11 := []int{1}
    var v12 uint32 = 2147483647
    var v13 int16 = -32768
    fmt.Println(v12, v12)
    _ = v11
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    var v1 int64 = 803
    var v2 string = "hello"
    v3 := []int{72}
    var v4 uint16 = 255
    var v5 int = -2147483648
    var v6 string = v2
    v14 := GenS1{F0: 2}
    if ("" <= ("foo bar" + v2)) {
        switch "" {
        case "\n\t":
            var v15 uint32 = 42
            _ = v15
        }
    }
    fmt.Println(v1, v2, v5)
    if true {
        for fi16 := 0; fi16 < 13; fi16++ {
            if true {
                if len(v3) > 0 {
                    v3[0] = (fi16 & 128)
                }
                for fi17 := 0; fi17 < 4; fi17++ {
                    v1 += v1
                    var v18 string = "\n\t"
                    _ = len(v3)
                    _ = fi17
                    _ = v18
                }
            }
            _ = fi16
        }
        v1 = (^(v1))
    }
    if false {
        if true {
            v5 -= 42
        } else {
            var v19 uint16 = v4
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v3) > 0 {
                    v3[0] = v5
                }
                v5 -= (v5 & -2147483648)
            } else {
                var v20 uint = 1
                _ = v20
            }
            _ = v19
        }
    } else {
        fmt.Println(v2, v5)
        switch v5 {
        case 42:
            v3 = append(v3, (v5 * (^(v5))))
        case 2147483647:
            v6 = ("hello" + v2)
        default:
            if len(v3) > 0 {
                v3[0] = 128
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if len(v3) > 0 {
        v3[0] = (v5 - v5)
    }
    if ((!(0.1 >= 1.5)) || true) {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint16 = v4
            _ = fi21
            _ = v22
        }
        var v23 bool = false
        _ = v23
    } else {
        switch (v5 % (((-(v5))) | 1)) {
        case 1:
            var v24 int64 = v1
            _ = v24
        case 32767:
            switch v5 {
            case -155:
                v3 = append(v3, (v5 / ((v5) | 1)))
            }
        case -835:
            var v25 bool = false
            _ = v25
        default:
            v3 = append(v3, 127)
        }
        var v26 int64 = -922
        _ = v26
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
