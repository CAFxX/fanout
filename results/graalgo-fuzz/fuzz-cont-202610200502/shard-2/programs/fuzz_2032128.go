package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int8, p1 uint32) int64 {
    p0 = p0
    var v5 int8 = p0
    if true {
        var v6 uint64 = 481
        _ = v6
    } else {
        v7 := []int{32767, 4294967295, 305}
        switch 394 {
        case -9223372036854775808:
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: -1}
                _ = fi8
                _ = v9
            }
        default:
            v10 := GenS2{F0: 9223372036854775807, F1: p1}
            _ = v10
        }
        _ = v7
    }
    fmt.Println(p1)
    v11 := GenS3{F0: 861}
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "hello"
    v3 := make(map[uint32]bool)
    var v4 int64 = 42
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 uint16 = 65535
        v4 -= 258
        v2 = (v2 + (v2 + "\n\t"))
        _ = fi12
        _ = v13
    }
    switch (("\n\t" + "a") + ("foo bar" + "\n\t")) {
    case "foo bar":
        v2 = "\n\t"
    case "\n\t":
        var v14 uint = 512
        _ = v14
    }
    v2 = ("" + ("hello" + v2))
    switch ("" + v2) {
    case "\n\t":
        v15 := GenS2{F0: -365, F1: 2}
        _ = v15
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            if (!false) {
                for fi17 := 0; fi17 < 9; fi17++ {
                    if (!(false && true)) {
                        v4 = v4
                        v4 += v4
                    }
                    var v18 string = ("" + ("\n\t" + v2))
                    _ = fi17
                    _ = v18
                }
                if false {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 int8 = -1
                        var v21 int8 = v20
                        var v22 uint8 = 42
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 9; fi23++ {
                        fmt.Println(v4)
                        fmt.Println(v2, v2, fi23)
                        v24 := []int{}
                        if false {
                            switch -208 {
                            case 42:
                                v4 = (v4 * 42)
                            default:
                                v4 -= v4
                            }
                        } else {
                            switch fi23 {
                            case 4294967295:
                                v25 := GenS3{F0: 668}
                                _ = v25
                            default:
                                v26 := GenS4{F0: (((v2 + "foo bar") + v2) + (v2 + (v2 + ("" + (((v2 + v2) + ((v2 + (v2 + v2)) + v2)) + (v2 + "\n\t")))))), F1: 232}
                                _ = v26
                            }
                            fmt.Println(fi16, v2)
                            var v27 uint = 18446744073709551615
                            _ = v27
                        }
                        _ = fi23
                        _ = v24
                    }
                }
            }
            var v28 uint64 = 2
            delete(v3, 2147483647)
            _ = fi16
            _ = v28
        }
    }
    var v29 int16 = 1000
    v4 -= 0
    var v30 float32 = 0.1
    switch -9223372036854775808 {
    case 42:
        for fi31 := 0; fi31 < 11; fi31++ {
            var v32 int = 32767
            var v33 uint16 = 111
            _ = fi31
            _ = v32
            _ = v33
        }
    case -2147483648:
        v34 := v3[2]
        _ = v34
    case -9223372036854775808:
        if (114 > 850) {
            if true {
                fmt.Println(v29, v4)
                if true {
                    switch (v2 + v2) {
                    case "hello":
                        for fi35 := 0; fi35 < 2; fi35++ {
                            v2 = ""
                            _ = len(v1)
                            v36 := GenS5{F0: v2}
                            _ = fi35
                            _ = v36
                        }
                    case "\n\t":
                        var v37 uint64 = 782
                        _ = v37
                    }
                    v30 -= v30
                    if false {
                        var v38 int32 = 42
                        v1[4294967295] = true
                        _ = v38
                    }
                }
                if (true || (-2147483648 == 60)) {
                    var v39 uint8 = 75
                    _ = v39
                } else {
                    v40 := GenS6{F0: 966, F1: 679, F2: 2147483647}
                    _ = v40
                }
            }
            v41 := GenS7{F0: 0}
            _ = v41
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
