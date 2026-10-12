package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int64; F1 int; F2 uint16 }
type GenS5 struct { F0 int16 }
func genFn0() int32 {
    var v6 int32 = -2147483648
    v7 := []int{}
    switch (((("a" + "hello") + "foo bar") + (("" + "foo bar") + "a")) + ("" + ("foo bar" + ("" + "a")))) {
    case "a":
        for fi8 := 0; fi8 < 11; fi8++ {
            v6 = v6
            for fi9 := 0; fi9 < 15; fi9++ {
                if true {
                    var v10 int64 = -1
                    v6 = v6
                    var v11 uint = 599
                    _ = v10
                    _ = v11
                }
                _ = len(v7)
                v6 = (v6 & v6)
                _ = fi9
            }
            _ = fi8
        }
    case "\n\t":
        v12 := GenS1{F0: -1}
        _ = v12
    }
    var v13 float32 = 0.1
    v14 := GenS2{F0: 1, F1: 1, F2: 1}
    _ = v13
    _ = v14
    _ = v6
    _ = v7
    return (v6 << ((0) & 63))
}

func main() {
    var v1 int16 = 277
    var v2 uint16 = 610
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 uint8 = 128
    v1 = (v1 ^ int16(v5))
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 uint32 = 42
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 int16 = v1
            fmt.Println(fi17, v1)
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v2, v2, v2)
            _ = fi19
        }
        _ = fi15
        _ = v16
    }
    var v20 int16 = v1
    if false {
        var v21 int32 = 42
        var v22 int = -2147483648
        _ = len(v4)
        _ = v21
        _ = v22
    }
    var v23 bool = (128 >= (v5 - (v5 & (v5 ^ v5))))
    fmt.Println(v23, v5)
    v20 -= (v20 | v20)
    v23 = true
    var v24 uint64 = 42
    v23 = false
    v5 += v5
    for fi25 := 0; fi25 < 20; fi25++ {
        if (4294967295 != 65535) {
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 int8 = 100
                var v28 int8 = (v27 >> ((2) & 63))
                v29 := []int{}
                var v30 uint16 = 2
                var v31 uint16 = (v30 + (v30 % (((v30 + v30)) | 1)))
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
                _ = v30
                _ = v31
            }
            v5 = v5
            if v23 {
                fmt.Println(v5)
                if ((639 <= 0) && ((-9223372036854775808 > -868) && (!v23))) {
                    var v32 int16 = (-(v1))
                    _ = v32
                }
            }
        } else {
            if (!(3.14159 > (3.14159 * (-2.25 - (1e10 * 3.14159))))) {
                v3[(fi25 / ((906) | 1))] = ""
            } else {
                if true {
                    v5 += v5
                    var v33 string = "a"
                    _ = v33
                } else {
                    if (!v23) {
                        fmt.Println(v23, v24, v20)
                        v34 := GenS2{F0: v2, F1: (v2 + v2), F2: 2}
                        v20 = (v20 / (((v1 & v1)) | 1))
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 13; fi35++ {
                        v36 := GenS3{F0: 65535, F1: "\n\t"}
                        v37 := GenS4{F0: 774, F1: (fi25 ^ int(v20)), F2: (^(v2))}
                        v38 := []int{(fi35 | -998), fi25}
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v1 = int16(v2)
                        _ = fi39
                    }
                }
                v2 -= (^(v2))
            }
        }
        if v23 {
            if false {
                for fi40 := 0; fi40 < 11; fi40++ {
                    switch (((("\n\t" + (("hello" + (("foo bar" + "hello") + ("" + ""))) + ("a" + "foo bar"))) + "\n\t") + ("foo bar" + "a")) + "") {
                    case "":
                        var v41 int8 = 0
                        _ = v41
                    case "hello":
                        var v42 uint16 = v2
                        _ = v42
                    }
                    v20 -= (v20 | (v20 << ((1) & 63)))
                    switch (fi25 ^ fi25) {
                    case -9223372036854775808:
                        if (!(v5 < (v5 % (((v5 | 1)) | 1)))) {
                            var v43 int8 = -45
                            _ = v43
                        } else {
                            switch (-(fi40)) {
                            case -928:
                                if v23 {
                                    var v44 int = (fi40 | 2147483647)
                                    var v45 uint32 = 42
                                    _ = v44
                                    _ = v45
                                } else {
                                    fmt.Println(v24, v5)
                                    var v46 float32 = 0.1
                                    _ = v46
                                }
                            case 4294967295:
                                v2 -= (v2 + v2)
                            case 9223372036854775807:
                                switch ("" + ((("\n\t" + "\n\t") + ("\n\t" + "a")) + ("a" + "hello"))) {
                                case "hello":
                                    var v47 int32 = -690
                                    _ = v47
                                case "\n\t":
                                    v48 := GenS5{F0: (v1 ^ (v20 >> ((492) & 63)))}
                                    _ = v48
                                }
                            default:
                                _ = len(v3)
                            }
                            v1 = (^(v1))
                        }
                    case 127:
                        v49 := []int{}
                        _ = v49
                    case 32768:
                        var v50 uint8 = 0
                        _ = v50
                    }
                    _ = fi40
                }
                var v51 uint8 = v5
                _ = v51
            }
            var v52 uint8 = 196
            _ = v52
        } else {
            v53 := []int{fi25}
            v54 := GenS5{F0: (^(v1))}
            _ = v53
            _ = v54
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
