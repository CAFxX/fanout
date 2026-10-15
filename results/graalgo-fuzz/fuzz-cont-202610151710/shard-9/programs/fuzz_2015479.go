package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 1
    v3 := []uint32{42, 2147483648}
    if true {
        var v4 uint64 = 2
        v5 := GenS1{F0: 32767}
        _ = v4
        _ = v5
    }
    var v6 int64 = (v2 << ((18446744073709551615) & 63))
    if (!("foo bar" >= "")) {
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS1{F0: 255}
                _ = fi8
                _ = v9
            }
            v2 += 603
            _ = fi7
        }
        var v10 uint8 = 135
        if (!false) {
            v6 -= (v6 * (v6 | (^(v2))))
        }
        _ = v10
    }
    v11 := GenS2{F0: 42}
    if false {
        var v12 float32 = -2.25
        var v13 int64 = v2
        _ = v12
        _ = v13
    } else {
        v6 = (v6 << ((1) & 63))
        var v14 float64 = 1.5
        var v15 uint32 = 2147483647
        _ = v14
        _ = v15
    }
    var v16 int16 = -1
    v17 := GenS1{F0: 42}
    var v18 uint32 = 201
    var v19 uint = 1
    switch 4294967295 {
    case 104:
        var v20 int64 = v2
        _ = v20
    case 901:
        v3 = append(v3, v18)
    }
    if false {
        if false {
            v2 = (-(v6))
            if len(v3) > 0 {
                v3[0] = v18
            }
        }
        var v21 int8 = 100
        var v22 int64 = (v2 - v2)
        _ = v21
        _ = v22
    } else {
        if true {
            switch -1 {
            case 913:
                for fi23 := 0; fi23 < 5; fi23++ {
                    v3 = append(v3, (v18 / ((v18) | 1)))
                    v2 = v2
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if true {
                            switch "hello" {
                            case "hello":
                                var v25 int16 = 42
                                _ = v25
                            case "":
                                v2 += 139
                            case "\n\t":
                                if (!((0 <= 0) && (29 == 125))) {
                                    var v26 uint64 = 42
                                    _ = len(v1)
                                    _ = v26
                                }
                                var v27 float32 = 1.5
                                _ = v27
                            }
                        } else {
                            var v28 int16 = 1
                            v6 = (v2 - v2)
                            var v29 float64 = 1.5
                            _ = v28
                            _ = v29
                        }
                        fmt.Println(v16)
                        v19 -= 2
                        _ = fi24
                    }
                    _ = fi23
                }
            default:
                switch (("a" + "\n\t") + "") {
                case "a":
                    var v30 float32 = -2.25
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 13; fi31++ {
                        var v32 int8 = 0
                        _ = fi31
                        _ = v32
                    }
                }
            }
        }
        var v33 uint = (v19 % ((255) | 1))
        if (true || false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 int64 = v6
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 13; fi37++ {
            if false {
                fmt.Println(fi36)
            }
            _ = fi37
        }
        _ = fi36
    }
    if (!true) {
        var v38 uint64 = 881
        v39 := GenS2{F0: 738}
        _ = v38
        _ = v39
    }
    if (-24 < 2147483647) {
        for fi40 := 0; fi40 < 17; fi40++ {
            fmt.Println(v18)
            if (false && ((v19 != v19) && (2 > 65535))) {
                var v41 int64 = -824
                for fi42 := 0; fi42 < 1; fi42++ {
                    v43 := GenS2{F0: 65535}
                    _ = len(v3)
                    _ = fi42
                    _ = v43
                }
                var v44 uint32 = (v18 >> (((v19 * (v19 * (v19 | (v19 - (^(v19))))))) & 63))
                _ = v41
                _ = v44
            } else {
                if (((v19 * v19) == 544) || ((fi40 & -913) > -444)) {
                    for fi45 := 0; fi45 < 9; fi45++ {
                        var v46 int16 = (v16 + (-(v16)))
                        switch (fi45 * (fi40 ^ 0)) {
                        case 2147483647:
                            v46 = 1
                        case 905:
                            if ((v18 + (^(v18))) < (v18 + v18)) {
                                var v47 float32 = 1e10
                                var v48 uint32 = v18
                                switch ("a" + "hello") {
                                case "foo bar":
                                    v19 = (^(v19))
                                case "":
                                    v49 := GenS3{F0: 42}
                                    _ = v49
                                case "\n\t":
                                    fmt.Println(v2, fi40, fi45)
                                }
                                v16 = -32768
                                _ = v47
                                _ = v48
                            } else {
                                fmt.Println(fi45, fi40)
                                var v50 uint16 = 1
                                var v51 uint64 = 42
                                _ = v50
                                _ = v51
                            }
                        case 0:
                            _ = len(v3)
                            var v52 int64 = (v2 >> ((v19) & 63))
                            var v53 int = -218
                            var v54 int16 = v16
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        _ = fi45
                        _ = v46
                    }
                }
                var v55 uint16 = 0
                _ = v55
            }
            _ = fi40
        }
    }
    var v56 bool = false
    v1[v18] = 65535
    for fi57 := 0; fi57 < 14; fi57++ {
        for fi58 := 0; fi58 < 3; fi58++ {
            fmt.Println(v2)
            _ = fi58
        }
        var v59 uint64 = 42
        _ = fi57
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v56
    _ = v6
}
