package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint64, p1 int8, p2 uint8) int32 {
    switch ("a" + "") {
    case "hello":
        var v6 int8 = (p1 * (p1 + (p1 - p1)))
        _ = v6
    case "foo bar":
        var v7 uint32 = 2
        _ = v7
    }
    switch 1 {
    case 124:
        var v8 int16 = 42
        _ = v8
    default:
        var v9 int32 = -110
        _ = v9
    }
    var v10 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}
func genFn1(p0 uint, p1 string) string {
    v11 := []int{1}
    p0 = (^(p0))
    if (42 != p0) {
        var v12 uint64 = 1
        for fi13 := 0; fi13 < 3; fi13++ {
            if (fi13 <= (fi13 + (-(fi13)))) {
                p0 = (p0 & 1)
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                p1 = "a"
                var v15 uint32 = 65535
                var v16 uint32 = v15
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 uint8 = 207
            _ = fi13
            _ = v17
        }
        _ = v12
    }
    if true {
        if true {
            if false {
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS1{F0: -1, F1: (fi18 >> ((p0) & 63))}
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if true {
                            var v21 int32 = -808
                            v11 = append(v11, (fi20 % ((2147483647) | 1)))
                            _ = v21
                        } else {
                            var v22 string = p1
                            _ = v22
                        }
                        switch (p1 + (p1 + "a")) {
                        case "a":
                            var v23 uint16 = 2
                            _ = v23
                        case "":
                            v24 := GenS2{F0: "", F1: 0, F2: (^(p0))}
                            _ = v24
                        case "foo bar":
                            if (true && true) {
                                fmt.Println(fi18, p1, fi20)
                            } else {
                                p0 = p0
                                p0 = (p0 ^ 4294967295)
                                p1 = p1
                            }
                        }
                        var v25 string = p1
                        _ = fi20
                        _ = v25
                    }
                    _ = fi18
                    _ = v19
                }
            }
            p0 = 42
        }
    } else {
        p0 = p0
        p1 = "\n\t"
        var v26 uint = p0
        _ = v26
    }
    _ = p0
    _ = p1
    _ = v11
    return p1
}

func main() {
    var v1 string = "a"
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    var v4 int64 = -1
    v5 := make(map[uint32]string)
    switch v2 {
    case "hello":
        switch 381 {
        case 1:
            fmt.Println(v4)
        case -430:
            if false {
                fmt.Println(v2, v4, v4)
            }
        default:
            for fi27 := 0; fi27 < 9; fi27++ {
                switch (fi27 << ((255) & 63)) {
                case 32767:
                    v2 = v2
                case 9223372036854775807:
                    var v28 int64 = 42
                    _ = v28
                case 505:
                    if (!true) {
                        var v29 bool = true
                        var v30 bool = (!v29)
                        _ = v29
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 float32 = (-2.25 / 1.0)
                    if true {
                        var v33 uint64 = 527
                        var v34 uint64 = (v33 % ((18446744073709551615) | 1))
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                var v35 string = (v1 + (("" + "\n\t") + (v1 + "hello")))
                _ = fi27
                _ = v35
            }
        }
    default:
        if (32767 < 65535) {
            if true {
                fmt.Println(v1)
                if (505 == 42) {
                    var v36 float32 = (0.0 + float32(1e10))
                    _ = v36
                } else {
                    v37 := GenS2{F0: (v1 + "hello"), F1: 197, F2: 0}
                    _ = v37
                }
                _ = len(v5)
            }
            v38 := GenS3{F0: 1, F1: true}
            switch v2 {
            case "\n\t":
                v39 := GenS2{F0: (("\n\t" + v2) + (v2 + v2)), F1: 468, F2: 40}
                _ = v39
            }
            _ = v38
        }
    }
    fmt.Println(v2, v2)
    var v40 uint64 = 0
    switch -807 {
    case 892:
        v41 := GenS4{F0: 68}
        _ = v41
    default:
        var v42 uint = 0
        _ = v42
    }
    if false {
        var v43 uint = 0
        switch v1 {
        case "\n\t":
            if true {
                v40 -= (v40 * v40)
                v2 = "foo bar"
            }
        case "a":
            v2 = ((((v1 + ("hello" + ((v2 + "foo bar") + "a"))) + (v2 + v1)) + v1) + (v1 + "hello"))
        default:
            var v44 bool = (!false)
            _ = v44
        }
        _ = v43
    } else {
        v45 := GenS5{F0: -1, F1: 308}
        _ = v45
    }
    var v46 uint = 0
    fmt.Println(v46, v4)
    for fi47 := 0; fi47 < 7; fi47++ {
        if (141 == 1) {
            v2 = "\n\t"
            switch (fi47 * (fi47 >> ((v46) & 63))) {
            case -993:
                v40 = 567
            case 267:
                v48 := GenS4{F0: -1}
                _ = v48
            case -9223372036854775808:
                v4 -= 111
            default:
                var v49 uint16 = 65535
                _ = v49
            }
            for fi50 := 0; fi50 < 18; fi50++ {
                v46 += 717
                v51 := GenS4{F0: 127}
                var v52 int64 = (v4 | -1)
                v40 = v40
                _ = fi50
                _ = v51
                _ = v52
            }
        }
        v53 := []int{fi47, (fi47 - -2147483648), fi47, fi47}
        _ = fi47
        _ = v53
    }
    v46 -= (^(v46))
    var v54 uint64 = (^(v40))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v5
    _ = v54
}
