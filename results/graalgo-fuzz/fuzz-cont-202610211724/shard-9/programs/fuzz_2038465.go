package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint64; F1 uint16; F2 int }
type GenS7 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint32, p1 int8, p2 uint) uint64 {
    v4 := []int{197, -238, 4294967295}
    switch 895 {
    case -558:
        p0 -= 65535
    case 32768:
        var v5 int8 = (-(p1))
        _ = v5
    case 2147483647:
        p0 -= (p0 << ((p2) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 int, p1 string) bool {
    p0 = p0
    var v6 uint = 42
    v7 := GenS1{F0: (v6 - (v6 - (v6 + v6))), F1: 0}
    v8 := []int{p0, (p0 / (((p0 - p0)) | 1)), p0, p0}
    if false {
        v9 := GenS2{F0: v6, F1: 2}
        _ = v9
    } else {
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 0
    v3 := make(map[int]int)
    if (!true) {
        v10 := []int{877, -757, 42, 9223372036854775807}
        if ((9223372036854775807 == 0) && false) {
            var v11 uint8 = 0
            var v12 int32 = 2147483647
            _ = v11
            _ = v12
        } else {
            fmt.Println(v2)
            if (-2147483648 >= 32767) {
                v13 := GenS3{F0: 42, F1: false}
                _ = v13
            } else {
                if true {
                    v14 := GenS3{F0: 2, F1: false}
                    v15 := GenS4{F0: -128, F1: 255, F2: 1}
                    var v16 uint8 = 127
                    _ = v14
                    _ = v15
                    _ = v16
                }
                if (1.5 > (0.1 - 0.0)) {
                    v2 = 844
                    v17 := GenS5{F0: 65535, F1: 2, F2: 65535}
                    _ = v17
                }
                fmt.Println(v2)
            }
        }
        _ = v10
    }
    if ("hello" != "\n\t") {
        switch 0 {
        case 0:
            if (v2 <= v2) {
                var v18 float32 = (1.5 - 3.14159)
                _ = v18
            }
        case 4294967295:
            if false {
                var v19 int = 42
                _ = v19
            } else {
                v20 := []int{32767, 32767, int(v2)}
                if false {
                    if true {
                        var v21 int8 = 0
                        v22 := GenS5{F0: 1, F1: 870, F2: 42}
                        v23 := GenS4{F0: (v21 | v21), F1: 0, F2: 464}
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        _ = len(v1)
                    }
                }
                var v24 uint32 = (v2 >> ((4294967295) & 63))
                _ = v20
                _ = v24
            }
        default:
            fmt.Println(v2)
        }
    }
    fmt.Println(v2)
    if true {
        if false {
            v2 = (v2 ^ (v2 ^ (v2 / (((v2 | v2)) | 1))))
            v2 += 2
            v25 := []int{4294967295}
            _ = v25
        } else {
            v26 := v1[-293]
            _ = v26
        }
    } else {
        var v27 float64 = 1e10
        fmt.Println(v2, v27)
        v27 += v27
        _ = v27
    }
    var v28 int16 = 1000
    v2 -= (v2 & v2)
    v2 = (v2 * (v2 * 42))
    var v29 float32 = 0.1
    for fi30 := 0; fi30 < 15; fi30++ {
        switch (fi30 * 9223372036854775807) {
        case 42:
            v29 -= v29
        case -2147483648:
            v31 := GenS5{F0: 979, F1: 0, F2: 42}
            _ = v31
        case 4294967295:
            var v32 int = 4294967295
            _ = v32
        default:
            var v33 uint8 = 2
            _ = v33
        }
        var v34 bool = (!((^(v2)) == v2))
        for fi35 := 0; fi35 < 9; fi35++ {
            if v34 {
                v2 -= v2
                if ("hello" >= "foo bar") {
                    v28 += (v28 | v28)
                    if (v34 || v34) {
                        delete(v1, 0)
                        if v34 {
                            v36 := []int{(fi30 ^ 42), 42}
                            v34 = ("\n\t" == ("hello" + ""))
                            v28 = (v28 << ((2) & 63))
                            var v37 uint32 = 2147483647
                            _ = v36
                            _ = v37
                        } else {
                            var v38 uint8 = 2
                            switch ("hello" + ("foo bar" + "")) {
                            case "\n\t":
                                v28 = v28
                            }
                            _ = v38
                        }
                        switch ((("hello" + "") + "\n\t") + "a") {
                        case "\n\t":
                            v34 = false
                        default:
                            if (v34 && (936 >= 1)) {
                                v39 := []int{fi30}
                                v40 := GenS6{F0: 2, F1: 255, F2: (fi35 << ((65535) & 63))}
                                _ = v39
                                _ = v40
                            } else {
                                fmt.Println(v34, v34, v29)
                                v34 = v34
                                fmt.Println(v28)
                            }
                            if (-1 < 298) {
                                var v41 uint32 = (v2 & v2)
                                fmt.Println(v2, v28)
                                _ = v41
                            } else {
                                v28 -= v28
                            }
                        }
                    }
                    fmt.Println(v34, v28)
                    v34 = (4294967295 == 65535)
                    var v42 string = "\n\t"
                    _ = v42
                }
            }
            var v43 bool = v34
            var v44 uint8 = 235
            var v45 bool = v34
            v43 = (!(int64(v44) != 9223372036854775807))
            if v34 {
                v46 := GenS7{F0: uint32(v44), F1: genFn0(v2, 42, 42)}
                if (!true) {
                    switch "a" {
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            v47 := GenS6{F0: 1, F1: 65535, F2: (fi30 + fi35)}
                            _ = v47
                        }
                    case "hello":
                        if (!(v44 != 52)) {
                            v29 += (1.5 + v29)
                            v48 := []int{(-(fi35))}
                            _ = v48
                        }
                    default:
                        v49 := v3[(^(fi30))]
                        _ = v49
                    }
                } else {
                    var v50 string = ("" + ("hello" + "foo bar"))
                    v51 := GenS5{F0: uint16(v2), F1: 0, F2: 65535}
                    _ = v50
                    _ = v51
                }
                if false {
                    var v52 bool = (v28 <= (v28 & (v28 << ((18446744073709551615) & 63))))
                    v53 := GenS2{F0: 1, F1: v44}
                    _ = v52
                    _ = v53
                }
                _ = v46
            } else {
                var v54 bool = true
                _ = v54
            }
            _ = fi35
            _ = v43
            _ = v44
            _ = v45
        }
        _ = fi30
        _ = v34
    }
    if true {
        var v55 int16 = v28
        v55 += 1000
        _ = v55
    } else {
        v56 := GenS7{F0: (v2 << ((65535) & 63)), F1: 1}
        _ = v56
    }
    v57 := GenS7{F0: 2147483647, F1: 988}
    if (3.14159 <= 3.14159) {
        var v58 string = "foo bar"
        switch 0 {
        case 32768:
            for fi59 := 0; fi59 < 18; fi59++ {
                v2 = (v2 << ((626) & 63))
                fmt.Println(v28, v28, v2)
                v3[(^(fi59))] = fi59
                _ = fi59
            }
        case 801:
            v29 += v29
        case -657:
            v60 := GenS2{F0: 989, F1: 255}
            _ = v60
        }
        _ = v58
    } else {
        var v61 int16 = (^(v28))
        _ = v61
    }
    fmt.Println(v29, v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v57
}
