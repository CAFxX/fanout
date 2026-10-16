package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 int32) int32 {
    p0 = p0
    p1 -= -1
    var v6 string = (("a" + (("" + "foo bar") + ("hello" + "a"))) + ("" + "\n\t"))
    p0 = p0
    v7 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    v3 := []uint32{486}
    v4 := []uint32{2147483647, 369, 0, 0}
    v5 := make(map[uint32]bool)
    var v8 int16 = 517
    if false {
        v2[42] = 0
    }
    if true {
        var v9 uint64 = 413
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint8 = 0
            if false {
                switch "foo bar" {
                case "\n\t":
                    v4 = append(v4, 65535)
                }
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                if (1 >= v8) {
                    var v13 string = ("foo bar" + ("\n\t" + ""))
                    v14 := v2[(fi10 - (fi10 >> ((255) & 63)))]
                    v11 += (v11 % ((v11) | 1))
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS2{F0: v11}
                    fmt.Println(fi10)
                    var v16 float64 = 3.14159
                    _ = v15
                    _ = v16
                }
                if false {
                    var v17 uint = 167
                    var v18 uint32 = 2
                    _ = v17
                    _ = v18
                } else {
                    var v19 int = fi12
                    if true {
                        var v20 uint8 = v11
                        if (true || false) {
                            if len(v4) > 0 {
                                v4[0] = 42
                            }
                            var v21 bool = (-1 <= 1)
                            v22 := v2[32767]
                            if (1.5 < 0.1) {
                                v9 += (^(v9))
                            } else {
                                delete(v5, 0)
                            }
                            switch v22 {
                            case 275:
                                var v23 int = (fi12 - 4294967295)
                                _ = v23
                            case -910:
                                v19 = -915
                            default:
                                var v24 uint16 = 2
                                _ = v24
                            }
                            switch -2147483648 {
                            case 9223372036854775807:
                                _ = len(v3)
                            default:
                                var v25 int8 = 113
                                _ = v25
                            }
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(v9, fi10)
                            v3 = append(v3, 0)
                            v26 := GenS1{F0: 4294967295}
                            _ = v26
                        }
                        v8 += (v8 >> ((2) & 63))
                        _ = v20
                    }
                    _ = v19
                }
                _ = fi12
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    if (!(!(!(588 == 895)))) {
        v8 += v8
    }
    switch -7 {
    case -2147483648:
        if false {
            fmt.Println(v8)
        }
    case 128:
        v27 := GenS2{F0: 255}
        _ = v27
    }
    var v28 int = 115
    v28 += (v28 / (((v28 & v28)) | 1))
    fmt.Println(v28, v8)
    delete(v5, 1)
    for fi29 := 0; fi29 < 8; fi29++ {
        v28 -= fi29
        v30 := GenS2{F0: 128}
        _ = fi29
        _ = v30
    }
    switch (v28 & -9223372036854775808) {
    case 0:
        if (true && (1.5 >= 1.5)) {
            for fi31 := 0; fi31 < 19; fi31++ {
                delete(v5, 65535)
                _ = fi31
            }
        }
    }
    v28 += (v28 - -385)
    switch v28 {
    case -2147483648:
        var v32 float32 = 1.5
        _ = v32
    case -295:
        v33 := GenS2{F0: 2}
        _ = v33
    case 42:
        var v34 uint = 65535
        _ = v34
    }
    if false {
        for fi35 := 0; fi35 < 7; fi35++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi36 := 0; fi36 < 2; fi36++ {
                v1["hello"] = false
                var v37 bool = true
                fmt.Println(v8, v37, v28)
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        v8 -= v8
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
