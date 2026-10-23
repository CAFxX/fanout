package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint16; F1 string; F2 bool }
type GenS3 struct { F0 uint64; F1 int; F2 uint16 }
func genFn0(p0 int16, p1 uint32, p2 bool) string {
    for fi5 := 0; fi5 < 1; fi5++ {
        if (1 <= 127) {
            p1 = p1
        }
        p1 -= (p1 & p1)
        _ = fi5
    }
    if (4294967295 != 2) {
        var v6 float64 = 3.14159
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                fmt.Println(fi7, fi7, v6)
                p1 += (p1 % (((p1 * (p1 | 169))) | 1))
                var v9 bool = (18446744073709551615 != 42)
                v10 := GenS1{F0: 518, F1: v9}
                v11 := []int{fi7, fi8}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            switch (^(fi7)) {
            case 187:
                v12 := GenS1{F0: 470, F1: false}
                _ = v12
            }
            _ = fi7
        }
        p2 = (!(127 == 128))
        _ = v6
    } else {
        var v13 int64 = 1
        var v14 uint64 = 42
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 18446744073709551615, F1: p2}
    v16 := GenS1{F0: 0, F1: (42 > 42)}
    for fi17 := 0; fi17 < 11; fi17++ {
        p1 -= p1
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return (("\n\t" + ("a" + ("\n\t" + "hello"))) + "")
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"foo bar", ""}
    var v3 int8 = 100
    var v4 int = 32767
    for fi18 := 0; fi18 < 8; fi18++ {
        switch "a" {
        case "":
            v19 := GenS1{F0: 966, F1: (2147483648 != 65535)}
            _ = v19
        case "hello":
            var v20 uint8 = 255
            _ = v20
        }
        _ = fi18
    }
    v3 += -50
    if true {
        for fi21 := 0; fi21 < 3; fi21++ {
            fmt.Println(v3, v3)
            _ = fi21
        }
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v4 -= (^(fi22))
        _ = fi22
    }
    var v23 uint16 = 607
    switch ("" + "hello") {
    case "hello":
        fmt.Println(v3, v23)
    case "a":
        for fi24 := 0; fi24 < 16; fi24++ {
            v4 += v4
            var v25 uint = 143
            _ = fi24
            _ = v25
        }
    }
    switch -495 {
    case 9223372036854775807:
        if (true || (127 == 255)) {
            v23 += (v23 ^ (^(v23)))
            if (42 == 42) {
                var v26 int8 = 100
                if (255 <= (v23 << ((255) & 63))) {
                    if ((v3 % (((v26 ^ 0)) | 1)) == v26) {
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 int = fi27
                            for fi29 := 0; fi29 < 11; fi29++ {
                                var v30 uint64 = 4294967295
                                _ = fi29
                                _ = v30
                            }
                            v31 := GenS2{F0: 483, F1: "\n\t", F2: (1.5 == 3.14159)}
                            _ = fi27
                            _ = v28
                            _ = v31
                        }
                        var v32 uint8 = 127
                        v33 := GenS3{F0: uint64(v32), F1: (v4 & v4), F2: v23}
                        _ = v32
                        _ = v33
                    } else {
                        if false {
                            v26 = (v26 >> ((18446744073709551615) & 63))
                            switch (^(v4)) {
                            case -2147483648:
                                _ = len(v2)
                            case -9223372036854775808:
                                if len(v2) > 0 {
                                    v2[0] = ("" + "hello")
                                }
                            case 660:
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    switch fi34 {
                                    case -933:
                                        var v35 bool = true
                                        _ = v35
                                    case 9223372036854775807:
                                        for fi36 := 0; fi36 < 8; fi36++ {
                                            v23 += v23
                                            _ = fi36
                                        }
                                    case 1:
                                        v37 := GenS2{F0: (v23 | v23), F1: "\n\t", F2: true}
                                        _ = v37
                                    default:
                                        fmt.Println(v23)
                                    }
                                    var v38 uint32 = 194
                                    _ = fi34
                                    _ = v38
                                }
                            }
                            var v39 int64 = 580
                            _ = v39
                        }
                        v40 := v1[(^(v4))]
                        _ = v40
                    }
                }
                if (true && ((true || true) || (2147483647 <= 4294967295))) {
                    v23 -= (^(v23))
                    var v41 uint8 = 2
                    v42 := GenS2{F0: 142, F1: "hello", F2: true}
                    _ = v41
                    _ = v42
                }
                _ = v26
            }
        }
    }
    if ((!true) && (true || false)) {
        v43 := GenS1{F0: 0, F1: false}
        for fi44 := 0; fi44 < 9; fi44++ {
            var v45 uint = uint(v23)
            var v46 float64 = 1e10
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = v43
    } else {
        v23 -= (v23 | v23)
        var v47 uint16 = 0
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
