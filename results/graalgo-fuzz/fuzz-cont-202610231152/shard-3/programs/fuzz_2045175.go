package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32, p1 int) uint64 {
    var v7 int8 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: "hello", F1: (-(p1)), F2: -402}
        v10 := []int{(fi8 >> ((4294967295) & 63)), int(v7)}
        switch ("\n\t" + "") {
        case "hello":
            var v11 uint32 = (^(p0))
            _ = v11
        case "":
            p1 -= p1
        case "\n\t":
            var v12 float64 = 1.5
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if (259 < 42) {
        v13 := GenS1{F0: (("hello" + ("a" + "a")) + ""), F1: -1, F2: 42}
        v14 := []int{}
        fmt.Println(p1, v7, p0)
        _ = v13
        _ = v14
    }
    switch (-(p1)) {
    case 9223372036854775807:
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 uint16 = 2
            _ = fi15
            _ = v16
        }
    }
    if true {
        var v17 string = "a"
        v17 = ""
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    v3 := []string{"", "\n\t"}
    var v4 uint = 20
    var v5 bool = (-106 >= -1)
    var v6 uint64 = 4294967295
    for fi18 := 0; fi18 < 14; fi18++ {
        switch "foo bar" {
        case "foo bar":
            v19 := GenS1{F0: ("a" + "hello"), F1: (fi18 | -2147483648), F2: -1}
            _ = v19
        case "a":
            v20 := GenS1{F0: "\n\t", F1: (fi18 * (-(fi18))), F2: 42}
            _ = v20
        case "hello":
            v4 = (v4 % (((^(v4))) | 1))
        }
        var v21 uint = v4
        v5 = v5
        _ = fi18
        _ = v21
    }
    v5 = (-1 <= 0)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint16 = 1
        for fi24 := 0; fi24 < 15; fi24++ {
            v4 -= v4
            v23 += (^(v23))
            if ((!false) && v5) {
                var v25 uint16 = 329
                v26 := GenS1{F0: "hello", F1: fi22, F2: 42}
                var v27 uint8 = 243
                _ = v25
                _ = v26
                _ = v27
            } else {
                fmt.Println(v5, v6)
                var v28 int8 = -80
                fmt.Println(v23)
                _ = v28
            }
            switch ("hello" + "\n\t") {
            case "foo bar":
                switch (("a" + ((("foo bar" + "\n\t") + "") + "foo bar")) + ("" + ("hello" + "foo bar"))) {
                case "":
                    switch 127 {
                    case 42:
                        var v29 uint = v4
                        _ = v29
                    case -622:
                        _ = len(v2)
                    }
                case "hello":
                    var v30 int8 = -1
                    _ = v30
                default:
                    v31 := GenS1{F0: "hello", F1: 32768, F2: 385}
                    _ = v31
                }
            case "":
                v3 = append(v3, "foo bar")
            default:
                if v5 {
                    fmt.Println(v6, v6, v6)
                    if true {
                        var v32 int = (fi22 / ((fi22) | 1))
                        if false {
                            v2[65535] = 656
                            if v5 {
                                var v33 string = ((("\n\t" + "hello") + ("hello" + "a")) + ("foo bar" + ""))
                                var v34 uint16 = 490
                                _ = v33
                                _ = v34
                            }
                            v35 := GenS1{F0: "a", F1: fi24, F2: int16(v6)}
                            _ = v35
                        }
                        v5 = (42 < -109)
                        _ = v32
                    } else {
                        var v36 int64 = 42
                        if v5 {
                            var v37 float64 = (-2.25 / (0.1 + 1e10))
                            _ = v37
                        }
                        _ = v36
                    }
                    var v38 uint = v4
                    _ = v38
                } else {
                    var v39 uint32 = 2147483648
                    var v40 uint = (^(v4))
                    v23 += (v23 << (((v4 ^ v4)) & 63))
                    _ = v39
                    _ = v40
                }
            }
            _ = fi24
        }
        switch "" {
        case "hello":
            v6 = uint64(v4)
        }
        _ = fi22
        _ = v23
    }
    switch ("\n\t" + "a") {
    case "hello":
        v41 := GenS1{F0: "\n\t", F1: 32768, F2: 737}
        _ = v41
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        v43 := GenS2{F0: (^(v6)), F1: 65535, F2: v6}
        v4 += (v4 & (v4 | (v4 - (v4 + (v4 & (v4 << ((v4) & 63)))))))
        fmt.Println(v4, v5)
        _ = fi42
        _ = v43
    }
    v44 := GenS1{F0: ("\n\t" + "foo bar"), F1: 939, F2: 907}
    v5 = false
    if (!v5) {
        var v45 int16 = -1
        _ = v45
    } else {
        for fi46 := 0; fi46 < 17; fi46++ {
            v5 = (v5 && v5)
            if true {
                var v47 int16 = 1
                v5 = (v5 || v5)
                fmt.Println(v47, v5, v4)
                _ = v47
            } else {
                if (!v5) {
                    v5 = v5
                    v48 := GenS3{F0: (false || ((v4 % ((v4) | 1)) > (v4 & (v4 & v4)))), F1: 255}
                    _ = v48
                }
            }
            _ = fi46
        }
        var v49 int32 = -900
        var v50 uint64 = genFn0(2, -234)
        _ = v49
        _ = v50
    }
    fmt.Println(v4, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
