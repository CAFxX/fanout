package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() int {
    var v6 uint32 = 209
    fmt.Println(v6)
    v6 = (v6 ^ v6)
    var v7 int8 = 45
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = 42
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    var v5 uint32 = 278
    v5 = (v5 / (((v5 & v5)) | 1))
    v8 := []int{1, -9223372036854775808}
    if (127 > 32767) {
        var v9 uint = 4294967295
        _ = v9
    } else {
        v2 = (v2 & -1)
        for fi10 := 0; fi10 < 2; fi10++ {
            v4[fi10] = true
            _ = fi10
        }
        var v11 int16 = -32768
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int8 = -38
        _ = fi12
        _ = v13
    }
    if false {
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int32 = 1
            var v16 uint16 = 2
            for fi17 := 0; fi17 < 6; fi17++ {
                v16 -= v16
                _ = fi17
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        var v18 float64 = 1.5
        v5 = 1
        _ = v18
    }
    var v19 int = 127
    switch genFn0() {
    case 9223372036854775807:
        fmt.Println(v19)
    default:
        v20 := v1[(v5 % ((42) | 1))]
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        _ = len(v8)
        _ = fi21
    }
    v22 := GenS1{F0: 65535}
    for fi23 := 0; fi23 < 6; fi23++ {
        v5 = v5
        if (-125 < 0) {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 float64 = -2.25
                v26 := GenS1{F0: uint16(v5)}
                fmt.Println(v5, v25, v25)
                var v27 int32 = -1
                switch ("" + ("a" + ("\n\t" + ("\n\t" + "foo bar")))) {
                case "\n\t":
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                default:
                    v19 = fi23
                }
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
            var v28 uint16 = 2
            var v29 uint = 4294967295
            _ = v28
            _ = v29
        } else {
            v19 = (v19 % ((fi23) | 1))
            switch -9223372036854775808 {
            case 32767:
                switch "foo bar" {
                case "a":
                    switch "hello" {
                    case "foo bar":
                        var v30 int8 = 0
                        _ = v30
                    case "\n\t":
                        v5 += (v5 & (v5 - (^(v5))))
                    }
                default:
                    v19 = (fi23 % ((-1) | 1))
                }
            case -9223372036854775808:
                for fi31 := 0; fi31 < 5; fi31++ {
                    v19 -= 9223372036854775807
                    fmt.Println(v5, v5, v19)
                    fmt.Println(v19, v19)
                    var v32 int16 = -1
                    _ = fi31
                    _ = v32
                }
            case 0:
                for fi33 := 0; fi33 < 2; fi33++ {
                    var v34 uint16 = 2
                    if false {
                        if (!false) {
                            v2 = 1
                            var v35 int16 = -1
                            _ = v35
                        } else {
                            var v36 string = "hello"
                            v2 = -199
                            _ = v36
                        }
                    }
                    _ = fi33
                    _ = v34
                }
            default:
                for fi37 := 0; fi37 < 14; fi37++ {
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        v38 := GenS1{F0: 65535}
                        _ = v38
                    case "":
                        switch (^(fi37)) {
                        case 9223372036854775807:
                            v2 = (v2 - v2)
                        case 32768:
                            v2 += 42
                            fmt.Println(v19, fi37)
                        default:
                            v19 -= -619
                        }
                    default:
                        switch "hello" {
                        case "hello":
                            var v39 uint64 = 44
                            var v40 uint64 = (v39 + v39)
                            fmt.Println(v5, fi37, v39)
                            _ = v39
                            _ = v40
                        case "\n\t":
                            fmt.Println(v19, v19)
                        }
                    }
                    _ = fi37
                }
            }
            switch -176 {
            case 42:
                var v41 float32 = 1e10
                _ = v41
            case -810:
                var v42 int = 32767
                _ = v42
            }
        }
        v43 := GenS1{F0: 1}
        _ = fi23
        _ = v43
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v44 := GenS1{F0: 255}
    var v45 bool = true
    fmt.Println(v45)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v8
}
