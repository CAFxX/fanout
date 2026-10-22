package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }
func genFn0(p0 uint, p1 string, p2 bool) int8 {
    p2 = false
    switch (p1 + "a") {
    case "hello":
        v6 := GenS1{F0: 1, F1: (p0 << ((p0) & 63)), F2: 65535}
        _ = v6
    case "\n\t":
        p1 = "a"
    default:
        var v7 int8 = 126
        _ = v7
    }
    var v8 string = ("foo bar" + "a")
    var v9 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -106
}
func genFn1(p0 uint32, p1 uint16) bool {
    switch -2147483648 {
    case 32767:
        v10 := []int{1}
        _ = v10
    case 937:
        if ((("" + "foo bar") + ("\n\t" + "foo bar")) != "hello") {
            if true {
                v11 := []int{}
                _ = v11
            } else {
                var v12 uint16 = (p1 | 2)
                if false {
                    var v13 int = -530
                    fmt.Println(p0, v13, v12)
                    _ = v13
                }
                if (530 <= -1) {
                    if (true && (int8(p1) > 57)) {
                        switch ("hello" + ("hello" + "hello")) {
                        case "foo bar":
                            v14 := GenS2{F0: ("" + "foo bar")}
                            _ = v14
                        }
                        v15 := GenS3{F0: 1}
                        _ = v15
                    } else {
                        var v16 uint8 = 127
                        _ = v16
                    }
                } else {
                    fmt.Println(p0, p1)
                    v17 := []int{0, 290, 128}
                    _ = v17
                }
                _ = v12
            }
            switch 127 {
            case -864:
                v18 := []int{-1}
                _ = v18
            case 0:
                p1 -= 42
            case 127:
                if (false || false) {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        if false {
                            v20 := GenS3{F0: -529}
                            var v21 float32 = 0.0
                            _ = v20
                            _ = v21
                        } else {
                            p0 = 829
                        }
                        v22 := GenS3{F0: 682}
                        _ = fi19
                        _ = v22
                    }
                    p1 = p1
                }
            }
            switch ("" + (("hello" + "") + ("\n\t" + "a"))) {
            case "a":
                v23 := []int{215, 32767, 2147483647}
                _ = v23
            default:
                var v24 int8 = 100
                _ = v24
            }
        }
    default:
        v25 := []int{int(p0), 2147483647}
        _ = v25
    }
    if (!(true || false)) {
        var v26 float64 = 0.1
        fmt.Println(p1, p0, p0)
        for fi27 := 0; fi27 < 6; fi27++ {
            for fi28 := 0; fi28 < 13; fi28++ {
                fmt.Println(fi28)
                var v29 int = (fi27 ^ fi27)
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        _ = v26
    } else {
        p0 += p0
        p1 += 65535
        var v30 uint64 = 18446744073709551615
        _ = v30
    }
    switch -2147483648 {
    case 127:
        if (4294967295 != 32767) {
            var v31 int16 = -32768
            switch 127 {
            case -23:
                if (!(1 < 61)) {
                    var v32 int32 = 42
                    v32 += (v32 >> ((2) & 63))
                    _ = v32
                }
            default:
                v33 := []int{1, 886, 32768}
                _ = v33
            }
            _ = v31
        } else {
            for fi34 := 0; fi34 < 3; fi34++ {
                for fi35 := 0; fi35 < 13; fi35++ {
                    p1 += (p1 | p1)
                    var v36 string = ("foo bar" + "a")
                    v37 := []int{(fi35 >> ((42) & 63))}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                var v38 int16 = 42
                _ = fi34
                _ = v38
            }
            if false {
                v39 := []int{454}
                switch "a" {
                case "\n\t":
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(p1)
                    case "a":
                        p0 += p0
                    }
                default:
                    switch ("\n\t" + (("foo bar" + ("foo bar" + "a")) + ("hello" + ("hello" + ("" + "a"))))) {
                    case "\n\t":
                        if false {
                            v40 := GenS2{F0: ("a" + ("" + ""))}
                            v41 := GenS1{F0: 1, F1: 42, F2: p0}
                            var v42 float32 = (3.14159 - (float32((-2.25 - -2.25)) / (0.1 + (float32((-2.25 + 1e10)) - 1e10))))
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    case "a":
                        p0 = (p0 % (((p0 << ((1) & 63))) | 1))
                    case "hello":
                        for fi43 := 0; fi43 < 8; fi43++ {
                            var v44 uint8 = 155
                            fmt.Println(v44)
                            switch "a" {
                            case "\n\t":
                                v44 = 255
                            case "":
                                v44 = 255
                            }
                            _ = fi43
                            _ = v44
                        }
                    default:
                        v45 := GenS4{F0: 0}
                        _ = v45
                    }
                }
                _ = v39
            }
        }
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        var v47 bool = (!(417 > -217))
        var v48 int64 = -549
        _ = fi46
        _ = v47
        _ = v48
    }
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 uint64 = 0
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    var v4 float32 = (v2 * -2.25)
    var v5 int16 = 0
    if true {
        var v49 uint8 = 128
        for fi50 := 0; fi50 < 6; fi50++ {
            if true {
                v51 := v3[42]
                _ = v51
            }
            _ = fi50
        }
        v52 := GenS2{F0: "\n\t"}
        _ = v49
        _ = v52
    }
    var v53 int16 = int16(v1)
    if true {
        var v54 float32 = 1e10
        fmt.Println(v54, v2, v2)
        for fi55 := 0; fi55 < 10; fi55++ {
            var v56 uint16 = 2
            var v57 float64 = (0.0 / -2.25)
            var v58 int = (fi55 << ((42) & 63))
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
        _ = v54
    } else {
        for fi59 := 0; fi59 < 20; fi59++ {
            if true {
                for fi60 := 0; fi60 < 16; fi60++ {
                    if false {
                        var v61 uint16 = 65535
                        switch (fi59 >> ((255) & 63)) {
                        case -2147483648:
                            fmt.Println(fi60, fi60)
                        }
                        switch fi59 {
                        case 1:
                            fmt.Println(v1, fi60)
                        default:
                            var v62 uint16 = (v61 << ((797) & 63))
                            _ = v62
                        }
                        _ = v61
                    } else {
                        v5 = int16(v1)
                        v63 := GenS5{F0: -2147483648}
                        v5 = v5
                        v64 := GenS1{F0: 0, F1: 65535, F2: 4294967295}
                        _ = v63
                        _ = v64
                    }
                    _ = fi60
                }
                v65 := GenS5{F0: -921}
                _ = v65
            } else {
                v4 -= v4
                v66 := v3[42]
                _ = v66
            }
            var v67 uint64 = v1
            var v68 uint8 = 0
            _ = fi59
            _ = v67
            _ = v68
        }
        var v69 int16 = 42
        _ = v69
    }
    v5 = 42
    v5 -= (v5 - 1)
    for fi70 := 0; fi70 < 9; fi70++ {
        fmt.Println(v4)
        _ = fi70
    }
    var v71 uint64 = v1
    v3[42] = ((("\n\t" == "hello") && false) && false)
    switch "hello" {
    case "hello":
        fmt.Println(v5, v71)
    case "\n\t":
        v72 := []int{-9223372036854775808, 127, -2147483648}
        _ = v72
    default:
        var v73 uint16 = 1
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v71
}
