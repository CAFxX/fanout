package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32; F2 int64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0() uint8 {
    switch -904 {
    case 4294967295:
        if false {
            var v6 int16 = 0
            var v7 int8 = 127
            _ = v6
            _ = v7
        } else {
            var v8 string = "\n\t"
            var v9 uint32 = 690
            for fi10 := 0; fi10 < 1; fi10++ {
                if (177 == 1) {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        var v12 string = "foo bar"
                        _ = fi11
                        _ = v12
                    }
                    if true {
                        switch v8 {
                        case "a":
                            if false {
                                var v13 uint = 255
                                var v14 uint32 = (v9 << ((255) & 63))
                                _ = v13
                                _ = v14
                            }
                        case "hello":
                            var v15 uint8 = 2
                            _ = v15
                        default:
                            v16 := GenS1{F0: 225}
                            _ = v16
                        }
                        fmt.Println(fi10, v8, v9)
                    }
                }
                var v17 float32 = 0.1
                fmt.Println(v17)
                _ = fi10
                _ = v17
            }
            _ = v8
            _ = v9
        }
    case 571:
        v18 := []int{-1, 32767, 1, 32767}
        _ = v18
    }
    var v19 uint = 255
    fmt.Println(v19)
    _ = v19
    return 255
}

func main() {
    var v1 uint32 = 1
    var v2 uint = 2
    v3 := []string{"", "foo bar", "a", "a"}
    v4 := []string{"", ("a" + ("foo bar" + "hello")), "hello", "hello"}
    var v5 uint8 = 128
    var v20 int8 = -3
    for fi21 := 0; fi21 < 13; fi21++ {
        v2 = (v2 % (((v2 | (v2 ^ v2))) | 1))
        v20 -= v20
        var v22 int32 = -2147483648
        _ = fi21
        _ = v22
    }
    var v23 int8 = v20
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := GenS1{F0: 0}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        v1 = v1
        _ = fi26
    }
    fmt.Println(v5, v1, v23)
    var v27 float32 = 0.1
    for fi28 := 0; fi28 < 6; fi28++ {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 uint32 = (v1 | (v1 & v1))
            v5 = (v5 + v5)
            v27 = (float32(float64(((0.0 - v27) + (-2.25 + v27)))) - float32(0.1))
            _ = fi29
            _ = v30
        }
        if false {
            _ = len(v4)
        } else {
            fmt.Println(v2, v20, v2)
            v31 := GenS2{F0: false, F1: (v1 ^ (^(v1))), F2: 42}
            switch fi28 {
            case -1:
                for fi32 := 0; fi32 < 19; fi32++ {
                    v27 += float32((float64((v27 - float32(-2.25))) / -2.25))
                    _ = fi32
                }
            default:
                v33 := GenS3{F0: 32767, F1: 1}
                _ = v33
            }
            _ = v31
        }
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 string = "foo bar"
            var v36 bool = true
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi28
    }
    v2 = (v2 << ((655) & 63))
    switch 127 {
    case 1:
        switch "" {
        case "\n\t":
            for fi37 := 0; fi37 < 11; fi37++ {
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                v3 = append(v3, "a")
                _ = fi37
            }
        default:
            v38 := GenS3{F0: -620, F1: uint64(v1)}
            _ = v38
        }
    case -1:
        var v39 int32 = 0
        _ = v39
    default:
        switch -1 {
        case 4294967295:
            var v40 uint = v2
            _ = v40
        case -35:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 127:
            if (!(0 > -1)) {
                switch "\n\t" {
                case "\n\t":
                    v41 := GenS4{F0: 4294967295, F1: genFn0()}
                    _ = v41
                case "hello":
                    v5 -= (v5 | (v5 >> ((42) & 63)))
                case "a":
                    if false {
                        var v42 int32 = 2147483647
                        _ = v42
                    }
                }
            }
        }
    }
    var v43 bool = true
    v20 = (v20 + v23)
    var v44 int16 = 228
    if v43 {
        v27 = float32(0.0)
        var v45 int = -916
        _ = v45
    }
    for fi46 := 0; fi46 < 14; fi46++ {
        v47 := GenS1{F0: -594}
        var v48 int32 = 0
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
