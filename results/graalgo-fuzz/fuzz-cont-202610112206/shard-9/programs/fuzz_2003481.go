package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int) int16 {
    p0 -= p0
    switch "\n\t" {
    case "a":
        var v6 bool = true
        _ = v6
    case "hello":
        var v7 uint32 = 1
        _ = v7
    case "":
        p0 = p0
    }
    v8 := []int{p0, p0, (p0 & (p0 * (p0 | p0))), (p0 * 42)}
    if len(v8) > 0 {
        v8[0] = (p0 - p0)
    }
    _ = p0
    _ = v8
    return -32768
}
func genFn1(p0 uint8, p1 uint64, p2 uint32) int64 {
    v9 := GenS1{F0: p1, F1: 2}
    var v10 float64 = 0.1
    v11 := []int{-620}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{65535, 0, 4294967295}
    v2 := []string{}
    var v3 float64 = 0.0
    var v4 int8 = 0
    var v5 uint16 = 0
    fmt.Println(v3, v4, v5)
    var v12 int64 = 624
    v13 := GenS1{F0: 4294967295, F1: (v5 & v5)}
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v3 = v3
    if true {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint16 = v5
            v12 = v12
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            v12 = genFn1(255, 2, 2147483647)
            if (true && false) {
                for fi17 := 0; fi17 < 6; fi17++ {
                    v5 = (v5 ^ 549)
                    _ = fi17
                }
                v5 -= v5
            }
            var v18 int32 = 42
            _ = fi16
            _ = v18
        }
        if false {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    } else {
        if (1.5 == 0.1) {
            switch "foo bar" {
            case "\n\t":
                for fi19 := 0; fi19 < 10; fi19++ {
                    v20 := GenS2{F0: 255, F1: 2}
                    if ((962 >= 901) && false) {
                        for fi21 := 0; fi21 < 13; fi21++ {
                            v22 := GenS2{F0: 2, F1: 2147483648}
                            _ = fi21
                            _ = v22
                        }
                        v5 += (v5 << ((42) & 63))
                        switch "a" {
                        case "foo bar":
                            var v23 uint32 = 2147483647
                            _ = v23
                        }
                    } else {
                        var v24 int16 = -1
                        _ = v24
                    }
                    _ = fi19
                    _ = v20
                }
            case "a":
                if true {
                    fmt.Println(v3, v5, v12)
                }
            case "":
                var v25 uint64 = 849
                _ = v25
            }
        }
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v4 -= 115
        if (!true) {
            fmt.Println(v4)
            var v27 uint64 = 4294967295
            for fi28 := 0; fi28 < 15; fi28++ {
                if false {
                    var v29 uint16 = 30
                    v30 := GenS1{F0: v27, F1: (v5 | v29)}
                    var v31 string = ("" + (("" + "\n\t") + ""))
                    v2 = append(v2, v31)
                    _ = v29
                    _ = v30
                    _ = v31
                }
                v27 -= v27
                _ = fi28
            }
            _ = v27
        }
        _ = fi26
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v32 := GenS2{F0: 255, F1: 4294967295}
    v33 := GenS3{F0: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
