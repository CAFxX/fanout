package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int32; F1 bool; F2 int64 }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0() uint {
    var v4 uint8 = 127
    v4 -= (^(v4))
    fmt.Println(v4)
    v4 -= v4
    _ = v4
    return 255
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    var v3 uint8 = 214
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 bool = (((0 < -1) || true) || false)
        _ = fi5
        _ = v6
    }
    if true {
        v7 := GenS1{F0: v3, F1: 473, F2: 2147483647}
        _ = v7
    }
    v3 -= v3
    v8 := GenS1{F0: v3, F1: 995, F2: 65535}
    _ = len(v1)
    if false {
        v3 += (v3 + v3)
        if true {
            switch "hello" {
            case "":
                var v9 uint32 = 4294967295
                _ = v9
            case "\n\t":
                fmt.Println(v3)
            case "foo bar":
                var v10 int32 = -1
                _ = v10
            default:
                if true {
                    v11 := GenS1{F0: 1, F1: 2, F2: 976}
                    var v12 bool = false
                    _ = v11
                    _ = v12
                } else {
                    var v13 int16 = 1000
                    for fi14 := 0; fi14 < 16; fi14++ {
                        fmt.Println(v13)
                        _ = fi14
                    }
                    var v15 string = ""
                    _ = v13
                    _ = v15
                }
            }
        } else {
            v3 -= v3
        }
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            v3 = 42
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS1{F0: (v3 * v3), F1: 913, F2: 4294967295}
                _ = fi17
                _ = v18
            }
            var v19 string = ""
            _ = fi16
            _ = v19
        }
        switch "a" {
        case "":
            v3 = (v3 ^ (v3 & v3))
        }
    }
    v20 := GenS2{F0: 2147483647, F1: 677}
    var v21 bool = true
    for fi22 := 0; fi22 < 8; fi22++ {
        switch 2147483647 {
        case 32768:
            switch -349 {
            case -2147483648:
                var v23 uint64 = 2
                _ = v23
            case 128:
                v24 := GenS1{F0: v3, F1: 32767, F2: 4294967295}
                _ = v24
            default:
                v3 -= (v3 / ((v3) | 1))
            }
        case 4294967295:
            fmt.Println(v3)
        default:
            v1["a"] = 968
        }
        v25 := GenS1{F0: v3, F1: 65535, F2: 4294967295}
        var v26 string = "a"
        _ = fi22
        _ = v25
        _ = v26
    }
    if v21 {
        v27 := GenS3{F0: 0, F1: (2 != 65535), F2: -413}
        if v21 {
            if false {
                v21 = true
            } else {
                v28 := GenS2{F0: -1, F1: 487}
                v29 := GenS2{F0: -2147483648, F1: 1}
                switch 2147483647 {
                case 32767:
                    v21 = (!true)
                case 0:
                    v30 := v1[""]
                    _ = v30
                }
                _ = v28
                _ = v29
            }
            v21 = (v21 && (false || (!(2 == 42))))
        } else {
            v3 -= (v3 ^ v3)
            var v31 int32 = 0
            _ = v31
        }
        v21 = v21
        _ = v27
    } else {
        if v21 {
            _ = len(v1)
            v32 := v1["foo bar"]
            _ = v32
        } else {
            v3 -= v3
            if v21 {
                if false {
                    v33 := GenS1{F0: 0, F1: 255, F2: 2}
                    v34 := GenS4{F0: 844, F1: 1}
                    var v35 int64 = 0
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    delete(v1, "hello")
                }
            } else {
                for fi36 := 0; fi36 < 10; fi36++ {
                    for fi37 := 0; fi37 < 9; fi37++ {
                        v21 = (1 == 255)
                        if len(v2) > 0 {
                            v2[0] = 641
                        }
                        var v38 int8 = -1
                        _ = fi37
                        _ = v38
                    }
                    var v39 int = -941
                    v3 = v3
                    _ = fi36
                    _ = v39
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v8
}
