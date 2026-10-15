package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint }
type GenS2 struct { F0 uint64; F1 int8; F2 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
func genFn0() bool {
    var v6 int8 = 127
    fmt.Println(v6)
    v6 += v6
    v7 := GenS1{F0: 9223372036854775807, F1: 0, F2: 966}
    var v8 int16 = 0
    _ = v6
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint8 = 255
    var v2 uint8 = (v1 | (v1 << ((255) & 63)))
    var v3 int = -10
    var v4 int16 = 32767
    v5 := make(map[string]int)
    v9 := GenS2{F0: 42, F1: 100, F2: v4}
    if (true && false) {
        if false {
            v10 := GenS1{F0: -598, F1: 2147483647, F2: 255}
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := v5["a"]
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 bool = true
                    _ = fi13
                    _ = v14
                }
                v15 := GenS3{F0: -2147483648}
                _ = fi11
                _ = v12
                _ = v15
            }
            for fi16 := 0; fi16 < 8; fi16++ {
                fmt.Println(v3, v2)
                v5[("\n\t" + "\n\t")] = (fi16 / (((v3 - fi16)) | 1))
                _ = fi16
            }
            _ = v10
        } else {
            fmt.Println(v3)
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := GenS4{F0: 107}
                    _ = fi18
                    _ = v19
                }
                fmt.Println(v2, v2)
                _ = fi17
            }
            switch "\n\t" {
            case "foo bar":
                var v20 bool = true
                _ = v20
            case "hello":
                var v21 uint64 = 1
                _ = v21
            }
        }
        switch "\n\t" {
        case "a":
            var v22 int16 = v4
            _ = v22
        case "hello":
            v23 := GenS3{F0: 1}
            _ = v23
        case "":
            var v24 float32 = 1e10
            _ = v24
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            v3 -= (^(fi25))
            _ = fi25
        }
    } else {
        v1 = v1
        var v26 int64 = -9223372036854775808
        for fi27 := 0; fi27 < 18; fi27++ {
            if true {
                if (0 >= 1) {
                    var v28 int = (v3 ^ -1)
                    var v29 string = "hello"
                    _ = v28
                    _ = v29
                }
            }
            _ = fi27
        }
        _ = v26
    }
    var v30 bool = true
    var v31 int32 = 503
    v32 := GenS2{F0: 821, F1: 64, F2: (^(v4))}
    v31 += -690
    if false {
        for fi33 := 0; fi33 < 8; fi33++ {
            fmt.Println(v30, v2, v30)
            v3 = fi33
            switch ("a" + "foo bar") {
            case "hello":
                var v34 uint16 = 42
                _ = v34
            case "a":
                if (4294967295 >= 2147483648) {
                    switch "a" {
                    case "foo bar":
                        v35 := GenS5{F0: v30}
                        _ = v35
                    default:
                        delete(v5, "")
                    }
                    v30 = (1 == 0)
                } else {
                    switch (v3 >> ((0) & 63)) {
                    case -2147483648:
                        for fi36 := 0; fi36 < 12; fi36++ {
                            delete(v5, "hello")
                            var v37 uint8 = 128
                            var v38 float32 = ((0.1 / float32(((float64((float32((1e10 * (0.0 - -2.25))) * 3.14159)) * 1e10) - 0.1))) * 3.14159)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    default:
                        v31 = (v31 * v31)
                    }
                    v2 = 128
                }
            default:
                fmt.Println(v4, v3, v2)
            }
            _ = fi33
        }
        var v39 string = "\n\t"
        v39 = "hello"
        _ = v39
    }
    v40 := []int{(-(v3)), int(v4)}
    v5["\n\t"] = -9223372036854775808
    var v41 bool = v30
    if v41 {
        v42 := GenS3{F0: -106}
        _ = v42
    }
    v2 = 42
    switch v3 {
    case -1:
        v1 += v1
    case 127:
        var v43 string = "a"
        _ = v43
    }
    switch -858 {
    case -9223372036854775808:
        v40 = append(v40, (v3 / ((128) | 1)))
    case 128:
        if false {
            v30 = v30
            _ = len(v40)
            switch "hello" {
            case "a":
                var v44 float64 = (0.0 + 0.0)
                _ = v44
            case "hello":
                v45 := GenS2{F0: 1, F1: 100, F2: v4}
                _ = v45
            }
        }
    }
    v46 := GenS5{F0: (161 > (v1 ^ (v1 / ((v1) | 1))))}
    fmt.Println(v41, v4, v41)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v5
    _ = v9
}
