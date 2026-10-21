package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int, p1 uint) uint32 {
    v5 := GenS1{F0: (!(127 == 1))}
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p0)
    case "a":
        if true {
            p0 = p0
            fmt.Println(p0)
        }
    }
    p1 = (^(p1))
    switch ((((("foo bar" + "") + (("hello" + "\n\t") + ("a" + ("foo bar" + "\n\t")))) + "hello") + "") + "") {
    case "a":
        v6 := GenS1{F0: true}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 327
}

func main() {
    v1 := []uint32{65535, 1, 0, 948}
    var v2 float64 = 1.5
    var v3 int16 = 42
    var v4 uint8 = 255
    v7 := GenS2{F0: 2147483648}
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint = 65535
        _ = fi8
        _ = v9
    }
    _ = len(v1)
    if false {
        var v10 string = "\n\t"
        _ = v10
    }
    if true {
        v4 = 128
        for fi11 := 0; fi11 < 17; fi11++ {
            v2 = 1e10
            v3 = 1000
            var v12 int = (fi11 ^ fi11)
            _ = fi11
            _ = v12
        }
        switch "hello" {
        case "hello":
            var v13 int32 = 0
            _ = v13
        case "foo bar":
            v4 = v4
        case "":
            if (3.14159 > 0.1) {
                var v14 bool = false
                _ = v14
            }
        default:
            var v15 uint32 = 1
            _ = v15
        }
    } else {
        var v16 string = (("" + ((("\n\t" + "hello") + "hello") + "foo bar")) + (("a" + ("" + (("foo bar" + (("hello" + "") + ("\n\t" + "a"))) + "hello"))) + "a"))
        var v17 string = ""
        _ = v16
        _ = v17
    }
    switch -1 {
    case 42:
        fmt.Println(v4)
    case 697:
        if (!false) {
            var v18 int64 = 667
            _ = v18
        } else {
            v2 = float64(3.14159)
            var v19 uint64 = 18446744073709551615
            if len(v1) > 0 {
                v1[0] = 0
            }
            _ = v19
        }
    case 849:
        v20 := GenS2{F0: 2147483647}
        _ = v20
    default:
        fmt.Println(v3, v2, v3)
    }
    var v21 uint8 = v4
    var v22 int8 = -1
    if (!false) {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 float64 = float64((1e10 / (0.1 * 3.14159)))
            v25 := GenS3{F0: 1}
            _ = len(v1)
            _ = fi23
            _ = v24
            _ = v25
        }
        v1 = append(v1, 2147483648)
    }
    if (299 == 2147483647) {
        var v26 int16 = v3
        v21 = (v4 | v21)
        _ = v26
    }
    if ((("a" + "a") >= "foo bar") || false) {
        switch (("a" + "a") + "hello") {
        case "a":
            if ((!(367 == 4294967295)) || (2 < 42)) {
                for fi27 := 0; fi27 < 2; fi27++ {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v4 += 127
                        switch (("" + "\n\t") + "a") {
                        case "foo bar":
                            v1 = append(v1, 4294967295)
                        case "a":
                            v4 = (v21 + 210)
                        }
                        var v29 bool = false
                        _ = fi28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 7; fi30++ {
                        if (true && false) {
                            var v31 int = (fi27 | fi27)
                            _ = len(v1)
                            fmt.Println(v2, v22)
                            if (((-32768 > (v3 - (v3 >> ((604) & 63)))) && (1 < 4294967295)) && false) {
                                fmt.Println(v2, v4)
                                v32 := GenS4{F0: (v3 == 0), F1: (v3 / ((v3) | 1)), F2: 64}
                                var v33 uint64 = 42
                                switch ("a" + ("\n\t" + "hello")) {
                                case "foo bar":
                                    v34 := GenS3{F0: 42}
                                    _ = v34
                                case "":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    if false {
                                        var v35 uint32 = genFn0(fi30, 255)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v35
                                    }
                                }
                                v36 := GenS2{F0: 1}
                                _ = v32
                                _ = v33
                                _ = v36
                            } else {
                                var v37 uint16 = 42
                                _ = v37
                            }
                            _ = v31
                        }
                        _ = fi30
                    }
                    _ = fi27
                }
            }
        case "foo bar":
            switch -9223372036854775808 {
            case 2147483647:
                switch "hello" {
                case "foo bar":
                    var v38 int16 = (v3 >> ((389) & 63))
                    _ = v38
                case "":
                    var v39 int32 = 0
                    _ = v39
                case "\n\t":
                    v2 -= 0.0
                default:
                    var v40 bool = ((!false) && (v21 >= v21))
                    _ = v40
                }
            case -1:
                var v41 uint64 = 42
                _ = v41
            default:
                var v42 float32 = 0.1
                _ = v42
            }
        }
        v4 += v21
    }
    v2 = v2
    for fi43 := 0; fi43 < 14; fi43++ {
        for fi44 := 0; fi44 < 17; fi44++ {
            v45 := GenS4{F0: ((v4 + 127) != 164), F1: (v3 % (((v3 / ((694) | 1))) | 1)), F2: 1}
            v46 := GenS5{F0: 618}
            _ = fi44
            _ = v45
            _ = v46
        }
        fmt.Println(fi43, v4)
        if true {
            for fi47 := 0; fi47 < 10; fi47++ {
                v48 := GenS2{F0: 632}
                var v49 string = "hello"
                var v50 int32 = 1
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
        }
        _ = fi43
    }
    _ = len(v1)
    if false {
        v2 = (float64((3.14159 / 3.14159)) / -2.25)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
}
