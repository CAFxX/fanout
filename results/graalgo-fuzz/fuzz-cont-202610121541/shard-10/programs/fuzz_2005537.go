package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint }

func main() {
    var v1 uint = 2
    v2 := []uint32{0}
    var v3 string = ""
    switch 32767 {
    case 32767:
        var v4 bool = false
        _ = v4
    case 4294967295:
        var v5 int16 = -228
        _ = v5
    }
    v6 := GenS1{F0: 2147483647}
    fmt.Println(v1, v1)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS1{F0: 42}
            _ = fi8
            _ = v9
        }
        v1 = v1
        fmt.Println(fi7)
        _ = fi7
    }
    var v10 float32 = (float32(1.5) * 1e10)
    v11 := GenS1{F0: -2147483648}
    var v12 uint8 = 254
    v3 = v3
    switch -1 {
    case 4294967295:
        if false {
            if (v10 != (((v10 - float32(0.0)) + v10) / -2.25)) {
                if (((-2.25 + 1e10) == -2.25) || (v12 > v12)) {
                    v13 := GenS1{F0: 1}
                    v14 := GenS2{F0: 2, F1: 4294967295}
                    var v15 uint16 = 201
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    if true {
                        for fi16 := 0; fi16 < 12; fi16++ {
                            if false {
                                for fi17 := 0; fi17 < 1; fi17++ {
                                    _ = len(v2)
                                    v18 := GenS3{F0: v1, F1: 962}
                                    _ = fi17
                                    _ = v18
                                }
                                var v19 uint = 2
                                _ = v19
                            }
                            for fi20 := 0; fi20 < 8; fi20++ {
                                v21 := GenS3{F0: (v1 | v1), F1: 42}
                                fmt.Println(v12)
                                _ = fi20
                                _ = v21
                            }
                            v22 := GenS3{F0: 2, F1: 768}
                            _ = fi16
                            _ = v22
                        }
                        v23 := GenS4{F0: 1, F1: 127}
                        var v24 int16 = 0
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 9; fi25++ {
                        fmt.Println(v10, v12, v12)
                        var v26 string = ""
                        v10 -= (v10 - 1e10)
                        _ = fi25
                        _ = v26
                    }
                    _ = len(v2)
                }
                for fi27 := 0; fi27 < 6; fi27++ {
                    if true {
                        var v28 bool = true
                        v29 := GenS5{F0: ("hello" + "\n\t")}
                        _ = v28
                        _ = v29
                    }
                    fmt.Println(v10, v10)
                    if true {
                        fmt.Println(v1, v3)
                    } else {
                        v30 := GenS6{F0: v1, F1: v1}
                        var v31 bool = false
                        fmt.Println(v3, v12)
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                }
                if false {
                    v3 = (v3 + "")
                    v1 = 169
                }
            } else {
                var v32 int8 = 31
                _ = v32
            }
            var v33 int16 = 618
            _ = v33
        } else {
            var v34 int8 = -120
            var v35 int8 = v34
            _ = v34
            _ = v35
        }
    case 447:
        var v36 bool = (!true)
        _ = v36
    default:
        var v37 float32 = v10
        _ = v37
    }
    var v38 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v6
}
