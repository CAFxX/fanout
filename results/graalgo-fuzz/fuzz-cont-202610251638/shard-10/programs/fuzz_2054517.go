package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 bool }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint64; F1 uint8; F2 string }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 string }
func genFn0(p0 uint, p1 uint64) int16 {
    v5 := []int{1, -708, 32768, 32767}
    if len(v5) > 0 {
        v5[0] = -58
    }
    v6 := GenS1{F0: 2}
    if (!false) {
        var v7 uint64 = 42
        if (p0 == p0) {
            var v8 uint8 = 1
            _ = v8
        } else {
            v9 := GenS1{F0: 0}
            v10 := GenS2{F0: 1, F1: v7, F2: (1 <= 2147483648)}
            _ = v10
            _ = v9
        }
        if len(v5) > 0 {
            v5[0] = -546
        }
        _ = v7
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(p1, p0, p0)
        v12 := GenS1{F0: 2}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 725
}
func genFn1(p0 uint8, p1 uint16) int16 {
    var v13 int64 = -1
    v14 := []int{127, int(v13)}
    p0 = p0
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := []uint32{405, 1, 2147483647, 2}
    v2 := make(map[int]uint32)
    var v3 string = "hello"
    var v4 int64 = 0
    v15 := GenS3{F0: "\n\t", F1: v3}
    switch (v3 + v3) {
    case "foo bar":
        switch ("foo bar" + ("foo bar" + (((v3 + "a") + "a") + v3))) {
        case "a":
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(fi16, v4, v3)
                v17 := GenS2{F0: 0, F1: 1, F2: true}
                _ = fi16
                _ = v17
            }
        case "\n\t":
            var v18 string = "a"
            _ = v18
        default:
            fmt.Println(v3)
        }
    }
    if false {
        switch ("a" + (((v3 + "a") + "a") + "hello")) {
        case "a":
            v4 = (v4 * 685)
        }
        switch 905 {
        case -2147483648:
            switch 42 {
            case 9223372036854775807:
                var v19 int32 = 42
                _ = v19
            }
        case -1:
            v4 = v4
        default:
            v3 = v3
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2[772] = 2147483648
        if (!(!(2 < 42))) {
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS4{F0: 2, F1: 127, F2: (v3 + (v3 + ""))}
                v4 = 217
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: 32767}
            _ = v22
        } else {
            _ = len(v1)
        }
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := GenS5{F0: 187, F1: 0}
            fmt.Println(fi23)
            if (!(false && ((255 == 13) && (!true)))) {
                var v26 string = (v3 + v3)
                if true {
                    v3 = ("hello" + "\n\t")
                    v3 = v3
                    var v27 int = fi23
                    _ = v27
                } else {
                    var v28 float32 = 1.5
                    _ = v28
                }
                delete(v2, (fi23 << ((479) & 63)))
                v26 = ("\n\t" + v3)
                switch (v26 + (((("foo bar" + v26) + ((v3 + (v26 + v3)) + v26)) + "foo bar") + "foo bar")) {
                case "a":
                    var v29 uint32 = 404
                    _ = v29
                case "":
                    v4 -= v4
                case "foo bar":
                    v30 := GenS3{F0: (("a" + "foo bar") + "hello"), F1: v3}
                    _ = v30
                }
                _ = v26
            }
            _ = fi24
            _ = v25
        }
        var v31 uint64 = 700
        _ = len(v2)
        _ = fi23
        _ = v31
    }
    fmt.Println(v4, v4)
    fmt.Println(v3)
    if false {
        if len(v1) > 0 {
            v1[0] = 221
        }
        if false {
            switch v3 {
            case "a":
                v3 = ("hello" + ("foo bar" + (("\n\t" + "\n\t") + (v3 + v3))))
            case "\n\t":
                v2[-2147483648] = 756
            case "foo bar":
                delete(v2, 0)
            }
            for fi32 := 0; fi32 < 5; fi32++ {
                for fi33 := 0; fi33 < 2; fi33++ {
                    if (91 <= 255) {
                        v34 := GenS5{F0: 128, F1: 839}
                        switch (fi33 * fi33) {
                        case 589:
                            v3 = (v3 + "\n\t")
                        case 32767:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        if (0.0 != 0.0) {
                            v3 = v3
                        } else {
                            v35 := GenS6{F0: v3}
                            if (!false) {
                                var v36 int16 = -1
                                _ = v36
                            } else {
                                var v37 int8 = 1
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v37
                            }
                            var v38 float32 = (0.1 / 3.14159)
                            v38 = v38
                            v38 = v38
                            var v39 int32 = 0
                            fmt.Println(v39, v38)
                            if (18446744073709551615 != 42) {
                                v38 = float32(-2.25)
                            }
                            _ = v35
                            _ = v38
                            _ = v39
                        }
                        _ = v34
                    }
                    _ = fi33
                }
                v1 = append(v1, 1)
                var v40 int8 = 7
                _ = fi32
                _ = v40
            }
            if (!true) {
                v4 = 9223372036854775807
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 0
            }
            for fi41 := 0; fi41 < 1; fi41++ {
                _ = len(v2)
                _ = fi41
            }
        }
    }
    var v42 bool = true
    for fi43 := 0; fi43 < 7; fi43++ {
        v4 = (v4 >> ((65535) & 63))
        var v44 uint16 = 65535
        v2[(fi43 - -2147483648)] = 2147483648
        _ = fi43
        _ = v44
    }
    v1 = append(v1, 2)
    for fi45 := 0; fi45 < 6; fi45++ {
        for fi46 := 0; fi46 < 8; fi46++ {
            v4 = v4
            v47 := GenS3{F0: "a", F1: "\n\t"}
            _ = fi46
            _ = v47
        }
        if (!(!v42)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v48 uint32 = 1
            _ = v48
        } else {
            switch (fi45 >> ((2) & 63)) {
            case 1:
                v3 = v3
            }
            if (!v42) {
                fmt.Println(v4)
            }
        }
        _ = fi45
    }
    var v49 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
}
