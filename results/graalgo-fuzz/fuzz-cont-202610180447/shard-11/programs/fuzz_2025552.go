package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 string; F2 uint16 }
func genFn0(p0 bool, p1 bool, p2 int32) bool {
    v4 := []int{1}
    v4 = append(v4, 759)
    if p1 {
        if len(v4) > 0 {
            v4[0] = -694
        }
    } else {
        var v5 int = 2147483647
        if len(v4) > 0 {
            v4[0] = 1
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (1 != 32767)
}

func main() {
    var v1 string = "foo bar"
    var v2 string = "a"
    var v3 int = 127
    fmt.Println(v1, v2, v2)
    if true {
        v6 := []int{(v3 * (v3 % ((0) | 1))), (v3 >> ((942) & 63)), v3}
        if false {
            v2 = ""
            switch (v3 - v3) {
            case -506:
                if ((!false) && false) {
                    var v7 uint64 = 397
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v9 := GenS1{F0: 20, F1: v7, F2: 894}
                        for fi10 := 0; fi10 < 2; fi10++ {
                            var v11 bool = false
                            v6 = append(v6, (fi8 >> ((42) & 63)))
                            _ = fi10
                            _ = v11
                        }
                        switch v2 {
                        case "a":
                            var v12 int16 = -1
                            _ = v12
                        }
                        _ = fi8
                        _ = v9
                    }
                    _ = v7
                }
            default:
                var v13 bool = false
                _ = v13
            }
            v2 = v2
        } else {
            fmt.Println(v2)
            v14 := GenS2{F0: false}
            _ = v14
        }
        v1 = v1
        _ = v6
    }
    v15 := GenS1{F0: 1, F1: 18446744073709551615, F2: 42}
    var v16 uint8 = 127
    v2 = "\n\t"
    if true {
        for fi17 := 0; fi17 < 6; fi17++ {
            v1 = "\n\t"
            _ = fi17
        }
    }
    var v18 uint16 = 876
    for fi19 := 0; fi19 < 19; fi19++ {
        switch ((v1 + "hello") + "\n\t") {
        case "":
            v2 = ((v1 + ("hello" + v2)) + (v1 + v2))
        case "a":
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 uint32 = 588
                if ((648 >= 42) || true) {
                    v16 -= v16
                    v22 := []int{fi20, fi19, v3}
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
        case "\n\t":
            fmt.Println(v2, v18)
        }
        v2 = ("foo bar" + "")
        switch (("hello" + "\n\t") + v2) {
        case "":
            v23 := GenS3{F0: 0, F1: uint64(v16), F2: 65535}
            _ = v23
        default:
            v18 = (v18 | (^(v18)))
        }
        _ = fi19
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        if false {
            switch (-(v3)) {
            case -231:
                var v25 uint = 153
                _ = v25
            }
            for fi26 := 0; fi26 < 15; fi26++ {
                fmt.Println(v1, v18)
                _ = fi26
            }
            if true {
                var v27 int8 = -89
                _ = v27
            }
        } else {
            var v28 uint32 = 2
            fmt.Println(v2, v28)
            for fi29 := 0; fi29 < 20; fi29++ {
                v30 := []int{v3, fi29, 62}
                fmt.Println(v16)
                switch 1 {
                case -772:
                    switch (^(fi24)) {
                    case 128:
                        v1 = ""
                    case 1:
                        var v31 bool = false
                        _ = v31
                    }
                case -2147483648:
                    if ((false && false) || (!false)) {
                        var v32 uint8 = 109
                        _ = v32
                    } else {
                        fmt.Println(fi29, v1)
                        var v33 int32 = 2147483647
                        v16 += (^(v16))
                        var v34 uint = 1
                        if len(v30) > 0 {
                            v30[0] = v3
                        }
                        _ = v33
                        _ = v34
                    }
                case 4294967295:
                    v35 := GenS1{F0: (v16 | v16), F1: 1, F2: (^(v18))}
                    _ = v35
                default:
                    v28 = (v28 * (v28 / ((231) | 1)))
                }
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        v36 := GenS4{F0: 4294967295, F1: (v2 + (("a" + v1) + v2)), F2: v18}
        v37 := []int{(fi24 | v3)}
        _ = fi24
        _ = v36
        _ = v37
    }
    v38 := []int{}
    var v39 int = int(v18)
    switch v2 {
    case "hello":
        if len(v38) > 0 {
            v38[0] = 976
        }
    case "":
        var v40 int8 = -128
        _ = v40
    default:
        v16 = v16
    }
    v41 := GenS2{F0: true}
    if len(v38) > 0 {
        v38[0] = (-(v3))
    }
    switch v2 {
    case "a":
        fmt.Println(v3, v3, v2)
    case "hello":
        var v42 int = 430
        _ = v42
    case "foo bar":
        fmt.Println(v2, v16)
    default:
        switch 1 {
        case -419:
            fmt.Println(v1, v39)
        case 128:
            var v43 int16 = 932
            _ = v43
        case -657:
            if len(v38) > 0 {
                v38[0] = (^(v39))
            }
        }
    }
    var v44 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v41
    _ = v44
}
