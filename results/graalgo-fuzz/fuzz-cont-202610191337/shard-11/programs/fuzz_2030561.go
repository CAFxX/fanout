package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int32 }
func genFn0(p0 string, p1 int, p2 string) bool {
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            v7 := GenS1{F0: (p1 + fi6)}
            p0 = p2
            if (!false) {
                switch ((p0 + (((p0 + p0) + (p0 + (p0 + "foo bar"))) + "a")) + (p0 + p0)) {
                case "\n\t":
                    for fi8 := 0; fi8 < 16; fi8++ {
                        v9 := []int{fi8, (fi8 % (((p1 / ((fi8) | 1))) | 1)), 32768}
                        p2 = (p2 + p0)
                        _ = fi8
                        _ = v9
                    }
                default:
                    p1 = (p1 << ((4294967295) & 63))
                }
            } else {
                for fi10 := 0; fi10 < 17; fi10++ {
                    p2 = p2
                    _ = fi10
                }
                var v11 uint64 = 173
                p0 = "\n\t"
                _ = v11
            }
            _ = v7
        } else {
            v12 := []int{p1, (-(fi6)), (p1 << ((272) & 63))}
            var v13 int8 = 100
            _ = v12
            _ = v13
        }
        p1 = fi6
        p0 = ((p0 + p0) + ((p2 + "a") + p2))
        _ = fi6
    }
    v14 := GenS1{F0: -1}
    p0 = p2
    v15 := []int{159, p1, (p1 + p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return false
}

func main() {
    var v1 int = 0
    v2 := []string{"hello"}
    v3 := []int{(v1 & v1), (^(v1)), 93}
    v4 := make(map[uint32]int)
    var v5 uint8 = 128
    v16 := GenS1{F0: (v1 / ((-2147483648) | 1))}
    var v17 uint8 = v5
    fmt.Println(v5)
    var v18 int16 = -32768
    for fi19 := 0; fi19 < 18; fi19++ {
        v18 = v18
        if (v17 >= 127) {
            for fi20 := 0; fi20 < 19; fi20++ {
                v21 := v4[665]
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                _ = fi20
                _ = v21
            }
        }
        _ = fi19
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS1{F0: fi22}
        v5 -= v5
        v18 = (v18 ^ v18)
        _ = fi22
        _ = v23
    }
    v5 -= v5
    v1 -= 32767
    for fi24 := 0; fi24 < 5; fi24++ {
        v3 = append(v3, (-(v1)))
        _ = fi24
    }
    v5 = v17
    v17 += (v17 >> ((2) & 63))
    if (1 != -128) {
        var v25 uint16 = 255
        _ = v25
    } else {
        _ = len(v2)
        v26 := v4[114]
        v18 = v18
        _ = v26
    }
    v27 := GenS2{F0: uint16(v17), F1: 690}
    var v28 uint64 = 1
    switch "hello" {
    case "":
        var v29 float32 = 0.1
        _ = v29
    case "\n\t":
        if false {
            if true {
                if false {
                    v3 = append(v3, (-(v1)))
                    for fi30 := 0; fi30 < 16; fi30++ {
                        var v31 uint32 = 76
                        v2 = append(v2, (("a" + "") + "hello"))
                        _ = fi30
                        _ = v31
                    }
                    fmt.Println(v17, v18)
                }
                delete(v4, 2147483647)
            }
            v17 -= v5
            if true {
                var v32 bool = (!(!(v17 == v17)))
                var v33 float32 = 3.14159
                _ = v32
                _ = v33
            }
        } else {
            v34 := GenS2{F0: 32767, F1: 0}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v34
        }
    case "hello":
        if false {
            var v35 int64 = -252
            v1 = (v1 ^ 127)
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
