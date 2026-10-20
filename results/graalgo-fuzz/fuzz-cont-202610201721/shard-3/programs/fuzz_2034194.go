package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int16; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 42
    var v3 int8 = -101
    if false {
        v4 := GenS1{F0: 18446744073709551615, F1: (v3 << ((4294967295) & 63))}
        var v5 int32 = -2147483648
        _ = v4
        _ = v5
    } else {
        v2 += 2147483647
        if (false || true) {
            fmt.Println(v2)
        } else {
            var v6 uint8 = 2
            v7 := GenS2{F0: 65535}
            fmt.Println(v6, v2)
            _ = v6
            _ = v7
        }
    }
    v8 := GenS3{F0: 42}
    v3 = -8
    var v9 uint16 = 522
    if true {
        fmt.Println(v9, v9)
    } else {
        v2 += v2
    }
    var v10 uint32 = 1
    var v11 uint8 = 1
    v12 := GenS2{F0: (v10 * (v10 % ((v10) | 1)))}
    switch "a" {
    case "":
        var v13 int16 = 975
        _ = v13
    }
    var v14 uint64 = 522
    switch ((("a" + ("a" + "a")) + "\n\t") + (("hello" + "") + ("" + "hello"))) {
    case "":
        var v15 uint8 = 222
        _ = v15
    case "\n\t":
        switch v2 {
        case 4294967295:
            if (!((v9 - (^(v9))) >= v9)) {
                var v16 int8 = -128
                _ = v16
            } else {
                fmt.Println(v3)
                fmt.Println(v11, v3)
                var v17 uint8 = v11
                _ = v17
            }
        }
    case "a":
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := []int{32768, v2}
            switch 128 {
            case 246:
                if false {
                    v20 := GenS1{F0: (v14 + v14), F1: v3}
                    fmt.Println(v9)
                    _ = v20
                }
            case 4294967295:
                switch "a" {
                case "\n\t":
                    v21 := GenS1{F0: (v14 + uint64(v9)), F1: v3}
                    _ = v21
                case "hello":
                    v22 := GenS3{F0: v10}
                    _ = v22
                case "":
                    var v23 int32 = int32(v3)
                    _ = v23
                }
            case 467:
                var v24 int8 = 1
                _ = v24
            default:
                var v25 uint64 = v14
                _ = v25
            }
            if (841 > -896) {
                v26 := GenS4{F0: (v3 << ((4294967295) & 63))}
                _ = v26
            }
            _ = fi18
            _ = v19
        }
    }
    v14 = (^(v14))
    var v27 int8 = -128
    v14 = v14
    switch "a" {
    case "foo bar":
        v28 := []int{}
        _ = v28
    default:
        switch ("\n\t" + ("hello" + "a")) {
        case "\n\t":
            if (false && false) {
                var v29 uint16 = (^(v9))
                for fi30 := 0; fi30 < 18; fi30++ {
                    v31 := []int{fi30, (fi30 - v2), 590}
                    var v32 float64 = 0.1
                    v33 := GenS4{F0: v27}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v29
            } else {
                v34 := GenS1{F0: 2, F1: (v27 << ((919) & 63))}
                _ = v34
            }
        case "a":
            if false {
                v35 := GenS2{F0: (v10 >> ((42) & 63))}
                for fi36 := 0; fi36 < 8; fi36++ {
                    var v37 bool = (!((v3 + v3) != v27))
                    var v38 string = ""
                    if v37 {
                        if false {
                            v37 = (42 < 958)
                            v9 += v9
                            switch fi36 {
                            case 2147483647:
                                v2 = (fi36 >> ((2) & 63))
                            case 128:
                                if v37 {
                                    v39 := GenS4{F0: (v27 & (v3 / (((v3 % ((v3) | 1))) | 1)))}
                                    _ = v39
                                } else {
                                    var v40 float64 = (((0.1 * (0.1 * 1.5)) / 1.0) * 0.1)
                                    v1[(v10 * (^(v10)))] = ("a" + "a")
                                    _ = v40
                                }
                            }
                        } else {
                            _ = len(v1)
                            var v41 string = (v38 + (("a" + "foo bar") + "foo bar"))
                            v42 := GenS5{F0: (v10 << ((18446744073709551615) & 63)), F1: 0, F2: int32(v11)}
                            _ = v41
                            _ = v42
                        }
                    }
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = v35
            } else {
                var v43 uint32 = v10
                delete(v1, (v10 * (^(v10))))
                switch ("" + "") {
                case "foo bar":
                    v44 := GenS1{F0: uint64(v10), F1: v27}
                    _ = v44
                }
                _ = v43
            }
        default:
            var v45 uint = 2
            _ = v45
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v8
    _ = v9
}
