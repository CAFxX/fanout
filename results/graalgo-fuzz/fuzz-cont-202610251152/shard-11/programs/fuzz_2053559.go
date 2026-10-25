package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint64 }
func genFn0(p0 string, p1 uint64, p2 bool) int32 {
    p0 = p0
    v5 := GenS1{F0: 1}
    v6 := GenS1{F0: 2}
    for fi7 := 0; fi7 < 7; fi7++ {
        p1 = p1
        _ = fi7
    }
    v8 := GenS2{F0: 148}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v8
    return 2147483647
}
func genFn1(p0 int16) uint32 {
    for fi9 := 0; fi9 < 8; fi9++ {
        p0 = (^(p0))
        switch "\n\t" {
        case "foo bar":
            v10 := GenS1{F0: 255}
            _ = v10
        }
        v11 := GenS1{F0: 2}
        _ = fi9
        _ = v11
    }
    var v12 int32 = -2147483648
    if (-46 <= -9223372036854775808) {
        v13 := []int{}
        v12 = v12
        p0 = p0
        _ = v13
    }
    v14 := []int{32767}
    v12 -= -745
    _ = p0
    _ = v12
    _ = v14
    return 4294967295
}

func main() {
    v1 := []int{32768}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v4 := []uint32{2147483647, 229, 2147483648, 971}
    var v15 uint = 255
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        if true {
            fmt.Println(v15)
            delete(v2, 1)
        } else {
            v15 -= v15
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2[2] = genFn1(1000)
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := v2[4294967295]
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    fmt.Println(v15)
    _ = len(v1)
    var v19 int = 127
    var v20 float32 = float32(0.1)
    delete(v3, (("\n\t" + "\n\t") + "foo bar"))
    v21 := GenS2{F0: 677}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 int32 = -2147483648
    for fi23 := 0; fi23 < 6; fi23++ {
        v22 = v22
        var v24 int16 = -32768
        if true {
            v25 := GenS2{F0: genFn1(v24)}
            if false {
                var v26 bool = ((-2.25 + ((3.14159 * -2.25) - (-2.25 / (3.14159 + 1.5)))) < 0.1)
                v19 = v19
                _ = len(v2)
                _ = v26
            }
            _ = len(v2)
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v20, v20, v20)
    if true {
        if true {
            var v27 int32 = v22
            _ = v27
        } else {
            fmt.Println(v15, v15, v15)
            var v28 uint16 = 892
            for fi29 := 0; fi29 < 8; fi29++ {
                v19 += -23
                v20 -= (v20 - ((v20 - v20) * (0.1 / (float32(1e10) - v20))))
                for fi30 := 0; fi30 < 20; fi30++ {
                    switch "a" {
                    case "a":
                        v15 += v15
                    case "\n\t":
                        v3["hello"] = (("a" + "\n\t") + "\n\t")
                    }
                    v19 += fi30
                    var v31 uint32 = 0
                    if (!(v28 > v28)) {
                        v31 -= v31
                    }
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
            _ = v28
        }
        v32 := GenS2{F0: 2147483648}
        var v33 int = (v19 / ((4294967295) | 1))
        _ = v32
        _ = v33
    } else {
        if false {
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 int64 = -167
                _ = fi34
                _ = v35
            }
            var v36 uint16 = 649
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 int64 = -1
                if true {
                    if len(v1) > 0 {
                        v1[0] = v19
                    }
                    v19 -= v19
                }
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        v39 := GenS3{F0: genFn1(-821), F1: 48, F2: 1}
        switch v19 {
        case 4294967295:
            v40 := GenS3{F0: 42, F1: 1, F2: uint64(v15)}
            _ = v40
        }
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
