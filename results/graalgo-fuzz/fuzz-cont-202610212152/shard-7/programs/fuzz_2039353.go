package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint) int64 {
    fmt.Println(p0)
    v4 := GenS1{F0: 42, F1: 42}
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -224
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 0
    var v3 int8 = 1
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint64 = 2
        _ = fi5
        _ = v6
    }
    var v7 uint = 180
    var v8 uint16 = 287
    for fi9 := 0; fi9 < 11; fi9++ {
        switch (fi9 | (fi9 - 9223372036854775807)) {
        case 128:
            v2 = (v2 - v2)
        default:
            var v10 uint8 = 238
            _ = v10
        }
        switch -2147483648 {
        case 128:
            v3 = 42
        case 1:
            _ = len(v1)
        case 342:
            if false {
                v8 += v8
                var v11 float32 = (-2.25 / 3.14159)
                _ = v11
            } else {
                if (!true) {
                    v12 := GenS1{F0: 4294967295, F1: 375}
                    if ((("a" < (("" + ("\n\t" + "a")) + "foo bar")) && (true && (!(false && (!true))))) || false) {
                        if ("foo bar" > ("" + ((("" + "\n\t") + ("foo bar" + "")) + (("" + ("" + "")) + "foo bar")))) {
                            var v13 uint32 = 0
                            _ = v13
                        } else {
                            if true {
                                delete(v1, 221)
                            }
                        }
                        if true {
                            if (1 > -1) {
                                if false {
                                    fmt.Println(v3, v2)
                                    v2 += v2
                                    fmt.Println(v2, v2)
                                }
                                for fi14 := 0; fi14 < 6; fi14++ {
                                    var v15 int = fi9
                                    var v16 uint32 = 2147483648
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            }
                            var v17 int64 = genFn0((v7 - (v7 * 18446744073709551615)))
                            v18 := []int{fi9, int(v3), 846, 42}
                            _ = v17
                            _ = v18
                        } else {
                            v19 := GenS1{F0: 1, F1: -2147483648}
                            _ = v19
                        }
                        fmt.Println(v8, v3, v8)
                    }
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 uint = 0
                        var v22 int8 = 100
                        v8 = v8
                        var v23 uint32 = 38
                        if false {
                            if ((-(v2)) <= -32768) {
                                fmt.Println(v2, v7)
                                v3 -= v3
                            }
                            if (1 < 2147483647) {
                                var v24 uint = (v21 >> ((v7) & 63))
                                _ = v24
                            } else {
                                var v25 bool = (!false)
                                _ = v25
                            }
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = v12
                } else {
                    var v26 uint64 = 1
                    v27 := GenS2{F0: "\n\t", F1: 0}
                    _ = v26
                    _ = v27
                }
            }
        default:
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 string = (("\n\t" + "\n\t") + "foo bar")
                fmt.Println(v7, v29, fi9)
                _ = fi28
                _ = v29
            }
        }
        if (v3 >= (^(v3))) {
            fmt.Println(v8, v3)
        } else {
            if (fi9 >= (-(fi9))) {
                fmt.Println(fi9)
                v7 = 1
            }
            v7 = (v7 * (v7 & (v7 ^ v7)))
        }
        _ = fi9
    }
    var v30 uint8 = 1
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 uint64 = 90
        var v33 uint32 = 2147483648
        switch "\n\t" {
        case "a":
            v7 = (v7 & 0)
        case "":
            var v34 uint8 = 12
            _ = v34
        case "foo bar":
            v35 := GenS3{F0: (v2 + v2)}
            _ = v35
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v3 = (v3 + v3)
    v36 := GenS2{F0: "\n\t", F1: -549}
    switch "a" {
    case "a":
        var v37 float32 = 0.0
        _ = v37
    default:
        switch -983 {
        case -1:
            v38 := []int{447, 128}
            _ = v38
        case 128:
            if (633 >= -452) {
                v1[333] = ""
            }
        default:
            fmt.Println(v30)
        }
    }
    var v39 int32 = 2147483647
    if (!true) {
        var v40 bool = true
        var v41 int32 = (v39 / ((v39) | 1))
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v39
    _ = v7
    _ = v8
}
