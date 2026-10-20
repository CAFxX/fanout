package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 uint16 }
func genFn0(p0 string, p1 string) string {
    v6 := GenS1{F0: 576, F1: 32767, F2: 2}
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 = p0
        p0 = p1
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}
func genFn1(p0 bool, p1 int16) uint16 {
    p0 = (1.5 <= (0.0 / 1.0))
    p0 = true
    p1 = p1
    if p0 {
        var v8 uint32 = 814
        _ = v8
    } else {
        var v9 uint = 42
        for fi10 := 0; fi10 < 18; fi10++ {
            v11 := []int{(fi10 / (((fi10 >> ((v9) & 63))) | 1)), 0, 4294967295, fi10}
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    _ = p0
    _ = p1
    return 115
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    v3 := []uint32{349, 2}
    var v4 string = "\n\t"
    v5 := make(map[uint32]string)
    var v12 int = 32767
    if (("foo bar" + v4) >= v4) {
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS1{F0: 65535, F1: 255, F2: 18446744073709551615}
            _ = fi13
            _ = v14
        }
        v12 = 717
    }
    v15 := GenS2{F0: "a", F1: 1, F2: 65535}
    v4 = (((v4 + ((v4 + ("foo bar" + "foo bar")) + v4)) + ("" + (v4 + "foo bar"))) + (("hello" + (v4 + "\n\t")) + ""))
    if false {
        var v16 int = (-(v12))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v16
    }
    if true {
        v1 += (v1 & (v1 ^ 381))
    } else {
        v12 = v12
        for fi17 := 0; fi17 < 6; fi17++ {
            fmt.Println(v4)
            switch (fi17 - (v12 / (((fi17 * 918)) | 1))) {
            case 857:
                for fi18 := 0; fi18 < 3; fi18++ {
                    fmt.Println(v4, v12)
                    if true {
                        v19 := GenS1{F0: 32767, F1: 330, F2: 4294967295}
                        var v20 uint64 = 0
                        v5[0] = (("foo bar" + "a") + v4)
                        _ = v19
                        _ = v20
                    } else {
                        var v21 bool = (1 != 100)
                        _ = v21
                    }
                    v1 += v1
                    _ = fi18
                }
            }
            var v22 string = ""
            _ = fi17
            _ = v22
        }
    }
    var v23 uint8 = 2
    var v24 uint16 = 228
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
