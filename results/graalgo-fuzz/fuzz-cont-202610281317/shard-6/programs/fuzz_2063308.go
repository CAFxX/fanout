package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 int; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 bool }
func genFn0() int {
    v7 := GenS1{F0: ("foo bar" + ("" + ("\n\t" + ("a" + "foo bar")))), F1: (!true), F2: 128}
    if true {
        var v8 int = 534
        _ = v8
    } else {
        var v9 uint32 = 4294967295
        var v10 uint16 = 898
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        switch (fi11 << ((0) & 63)) {
        case -2147483648:
            var v12 uint16 = 720
            _ = v12
        }
        var v13 string = ""
        if (-253 < 1) {
            switch (fi11 & fi11) {
            case 9223372036854775807:
                var v14 int32 = 1
                _ = v14
            default:
                switch (-(fi11)) {
                case 127:
                    v15 := []int{4294967295, (fi11 / ((fi11) | 1))}
                    _ = v15
                default:
                    v16 := GenS1{F0: ("" + ("a" + ("" + "\n\t"))), F1: true, F2: (fi11 & (fi11 & (fi11 & 128)))}
                    _ = v16
                }
            }
            v17 := GenS1{F0: ("hello" + v13), F1: false, F2: fi11}
            var v18 uint64 = 963
            _ = v17
            _ = v18
        } else {
            var v19 uint64 = 453
            _ = v19
        }
        _ = fi11
        _ = v13
    }
    var v20 int = -353
    v21 := []int{v20, v20, (v20 ^ (v20 & (v20 + 4294967295))), v20}
    _ = v20
    _ = v21
    _ = v7
    return (v20 >> ((1) & 63))
}
func genFn1(p0 string, p1 int8) int16 {
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := GenS2{F0: fi22, F1: 594, F2: 1}
        p0 = p0
        var v24 uint32 = 0
        _ = fi22
        _ = v23
        _ = v24
    }
    if (false && true) {
        var v25 int32 = -1
        v25 = v25
        _ = v25
    } else {
        switch p0 {
        case "a":
            for fi26 := 0; fi26 < 14; fi26++ {
                p1 += p1
                p0 = p0
                _ = fi26
            }
        }
    }
    fmt.Println(p0, p1)
    for fi27 := 0; fi27 < 9; fi27++ {
        switch fi27 {
        case 0:
            v28 := GenS3{F0: 42, F1: 32767, F2: 196}
            _ = v28
        case 32767:
            var v29 bool = (578 > 1)
            _ = v29
        case 4294967295:
            var v30 uint16 = 1
            _ = v30
        }
        var v31 float32 = 1e10
        for fi32 := 0; fi32 < 17; fi32++ {
            p0 = ("hello" + p0)
            var v33 int64 = 1
            if ((true || true) || false) {
                v34 := []int{genFn0(), 32768}
                _ = v34
            }
            _ = fi32
            _ = v33
        }
        _ = fi27
        _ = v31
    }
    _ = p0
    _ = p1
    return -926
}

func main() {
    var v1 int16 = 42
    var v2 bool = false
    var v3 uint8 = 128
    v4 := make(map[uint32]string)
    v5 := []int{}
    var v6 uint32 = 2
    v3 = 42
    v1 -= (v1 ^ v1)
    var v35 uint64 = uint64(v3)
    for fi36 := 0; fi36 < 12; fi36++ {
        var v37 float64 = 3.14159
        _ = fi36
        _ = v37
    }
    if v2 {
        if v2 {
            var v38 uint8 = 42
            v2 = v2
            _ = v38
        }
        v1 = -1
    } else {
        v39 := GenS4{F0: uint32(v3), F1: (!(0 == v35)), F2: -32768}
        for fi40 := 0; fi40 < 4; fi40++ {
            var v41 uint16 = 350
            v1 += int16(v35)
            v42 := GenS4{F0: 1, F1: ((v41 / (((^(v41))) | 1)) >= v41), F2: (v1 ^ (-(v1)))}
            _ = fi40
            _ = v41
            _ = v42
        }
        _ = v39
    }
    if v2 {
        v6 += (^(v6))
        if false {
            var v43 bool = false
            switch 0 {
            case -291:
                for fi44 := 0; fi44 < 13; fi44++ {
                    v6 = v6
                    if (!(!v43)) {
                        var v45 int16 = v1
                        _ = v45
                    }
                    v43 = (167 > 32767)
                    _ = fi44
                }
            case 32767:
                _ = len(v5)
            case 32768:
                var v46 int16 = (v1 << ((4294967295) & 63))
                _ = v46
            default:
                for fi47 := 0; fi47 < 10; fi47++ {
                    fmt.Println(v43)
                    fmt.Println(v2, fi47)
                    _ = fi47
                }
            }
            var v48 string = "a"
            _ = v43
            _ = v48
        }
        var v49 uint8 = (v3 % (((v3 - v3)) | 1))
        _ = v49
    } else {
        var v50 float32 = 0.1
        _ = v50
    }
    v2 = v2
    v51 := GenS3{F0: 2, F1: uint16(v35), F2: (^(v3))}
    var v52 uint = 171
    if true {
        if v2 {
            v52 -= 255
        }
    }
    v52 = v52
    if (!(v1 > 0)) {
        if v2 {
            for fi53 := 0; fi53 < 20; fi53++ {
                if (false && (int16(v6) >= 42)) {
                    var v54 uint = uint(v3)
                    _ = v54
                } else {
                    var v55 int32 = -1
                    v6 = v6
                    for fi56 := 0; fi56 < 5; fi56++ {
                        v35 = (v35 | v35)
                        var v57 string = ("a" + "hello")
                        _ = fi56
                        _ = v57
                    }
                    _ = v55
                }
                _ = fi53
            }
            if ("a" > "") {
                if (235 != 42) {
                    var v58 int64 = 42
                    _ = v58
                } else {
                    v59 := GenS3{F0: v52, F1: 42, F2: 127}
                    var v60 uint = (^(v52))
                    v52 = v60
                    _ = v59
                    _ = v60
                }
                var v61 string = "a"
                _ = v61
            } else {
                if ((^(v52)) != v52) {
                    for fi62 := 0; fi62 < 7; fi62++ {
                        fmt.Println(v3, v52)
                        fmt.Println(v2)
                        _ = fi62
                    }
                    switch 2147483647 {
                    case 32767:
                        var v63 uint = 145
                        _ = v63
                    case 42:
                        v64 := v4[v6]
                        _ = v64
                    case 4294967295:
                        var v65 uint8 = 0
                        _ = v65
                    }
                }
            }
            v66 := GenS2{F0: 32767, F1: v6, F2: v3}
            _ = v66
        } else {
            v67 := GenS1{F0: "foo bar", F1: false, F2: 586}
            v68 := GenS5{F0: v2}
            _ = v67
            _ = v68
        }
        for fi69 := 0; fi69 < 13; fi69++ {
            if true {
                for fi70 := 0; fi70 < 4; fi70++ {
                    var v71 bool = true
                    v6 = 2
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi70
                    _ = v71
                }
                v52 = 201
            }
            v72 := GenS6{F0: v2}
            var v73 bool = (genFn1("", 0) < v1)
            _ = fi69
            _ = v72
            _ = v73
        }
    } else {
        v1 = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
