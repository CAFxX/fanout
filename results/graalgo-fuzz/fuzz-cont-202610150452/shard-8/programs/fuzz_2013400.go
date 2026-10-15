package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint; F2 int8 }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint16; F2 string }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 int64; F1 int16; F2 int64 }
type GenS9 struct { F0 uint64 }
func genFn0() int8 {
    switch "foo bar" {
    case "foo bar":
        v5 := GenS1{F0: (!true), F1: 1}
        _ = v5
    }
    switch "" {
    case "foo bar":
        v6 := GenS2{F0: -936, F1: 42}
        _ = v6
    case "":
        switch "foo bar" {
        case "\n\t":
            if true {
                var v7 int32 = 2147483647
                var v8 uint8 = 255
                _ = v7
                _ = v8
            } else {
                switch 443 {
                case -1:
                    v9 := []int{-312, -9223372036854775808, -1}
                    _ = v9
                }
            }
        case "foo bar":
            if (!(-759 < 32767)) {
                v10 := GenS2{F0: 1, F1: 2}
                var v11 uint8 = 2
                fmt.Println(v11)
                _ = v10
                _ = v11
            }
        }
    case "a":
        var v12 uint32 = 65535
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 string = "hello"
            var v15 int64 = 630
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 int16 = 1000
        v16 -= 835
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := []int{}
        if (!false) {
            v19 := GenS1{F0: (0.0 > 1.5), F1: -1}
            switch "" {
            case "":
                for fi20 := 0; fi20 < 6; fi20++ {
                    fmt.Println(fi17, fi17)
                    switch fi20 {
                    case 2147483647:
                        var v21 int64 = 480
                        _ = v21
                    case -158:
                        fmt.Println(fi20, fi20)
                    case -814:
                        var v22 uint16 = 0
                        _ = v22
                    }
                    _ = fi20
                }
            }
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    if ((705 <= 1) || (!(false && true))) {
        var v23 int32 = 384
        _ = v23
    } else {
        if true {
            if true {
                var v24 string = "a"
                v25 := GenS3{F0: 189, F1: 42, F2: -113}
                _ = v24
                _ = v25
            } else {
                if (!false) {
                    var v26 uint16 = 301
                    v27 := []int{598}
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS2{F0: 1, F1: 290}
                    _ = v28
                }
                var v29 uint64 = 437
                if (((-2147483648 < -1) && true) || false) {
                    fmt.Println(v29)
                    v29 += 1
                    v29 += 747
                }
                _ = v29
            }
        }
    }
    return 78
}
func genFn1(p0 string) int32 {
    var v30 bool = false
    v31 := GenS4{F0: 1, F1: v30, F2: 120}
    _ = p0
    _ = v30
    _ = v31
    return -1
}

func main() {
    var v1 int8 = 1
    v2 := []string{"a", ("hello" + "a")}
    v3 := make(map[uint32]bool)
    var v4 uint8 = 255
    v32 := GenS5{F0: -413, F1: 0}
    v33 := GenS6{F0: 42, F1: 65535, F2: ("a" + "")}
    var v34 float32 = 1.5
    if (true && (-32768 != 1000)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            v35 := GenS7{F0: 1, F1: false}
            _ = v35
        } else {
            for fi36 := 0; fi36 < 16; fi36++ {
                v4 = (v4 % ((127) | 1))
                if true {
                    v37 := GenS8{F0: 9223372036854775807, F1: 0, F2: 0}
                    if false {
                        var v38 int16 = 1
                        _ = v38
                    }
                    _ = v37
                } else {
                    v3[4294967295] = (65535 < 2)
                    v39 := GenS2{F0: -753, F1: 4294967295}
                    _ = v39
                }
                _ = fi36
            }
            var v40 uint = 612
            _ = v40
        }
    }
    fmt.Println(v1)
    if false {
        var v41 uint16 = 837
        fmt.Println(v34, v1, v41)
        _ = v41
    } else {
        fmt.Println(v1, v1, v34)
        var v42 float64 = (((0.0 * 0.1) * (1.5 * 3.14159)) * 1.5)
        _ = v42
    }
    var v43 int = 234
    v4 -= 116
    v4 = v4
    var v44 uint16 = 383
    var v45 bool = ("foo bar" >= "a")
    v43 = v43
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    if (1e10 >= 1e10) {
        for fi46 := 0; fi46 < 18; fi46++ {
            var v47 float32 = v34
            v48 := GenS9{F0: 422}
            _ = fi46
            _ = v47
            _ = v48
        }
        var v49 int = int(v44)
        _ = len(v2)
        _ = v49
    } else {
        for fi50 := 0; fi50 < 18; fi50++ {
            delete(v3, uint32(v44))
            _ = fi50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
