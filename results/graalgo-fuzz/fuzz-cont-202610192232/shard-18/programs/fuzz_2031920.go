package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint16; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 int16; F2 uint64 }

func main() {
    v1 := []uint32{2147483647, 4294967295}
    v2 := []string{}
    var v3 int16 = 32767
    v4 := make(map[int]uint32)
    v3 -= (^(v3))
    var v5 int = 32768
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
        if true {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: (fi6 << ((65535) & 63)), F1: v5, F2: fi6}
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 12; fi10++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi10
            }
            switch ((("a" + "") + "hello") + ((((("" + "\n\t") + "\n\t") + (("a" + "hello") + ("foo bar" + "a"))) + "a") + "a")) {
            case "":
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 uint32 = 450
                    var v13 uint8 = 128
                    v14 := GenS1{F0: v5, F1: (fi6 * (-(fi6))), F2: int(v3)}
                    fmt.Println(fi11, fi6, v3)
                    v12 += 511
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
        } else {
            for fi15 := 0; fi15 < 5; fi15++ {
                var v16 uint8 = 128
                _ = fi15
                _ = v16
            }
        }
        if false {
            v5 = fi6
        } else {
            fmt.Println(v3, v5)
            if (!(700 >= 190)) {
                var v17 int = v5
                if len(v2) > 0 {
                    v2[0] = ("a" + "")
                }
                v3 = (v3 * v3)
                _ = v17
            }
        }
        _ = fi6
    }
    v2 = append(v2, "a")
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 string = ("" + "hello")
            v3 = (v3 >> ((65535) & 63))
            _ = fi19
            _ = v20
        }
        v5 = fi18
        v5 -= (fi18 | (fi18 | fi18))
        _ = fi18
    }
    v21 := GenS2{F0: ("hello" + "a")}
    _ = len(v4)
    for fi22 := 0; fi22 < 20; fi22++ {
        v3 -= (v3 - (v3 % ((1000) | 1)))
        v23 := GenS2{F0: ("\n\t" + "foo bar")}
        v24 := GenS3{F0: true, F1: 779, F2: -118}
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 int16 = v3
        v27 := GenS1{F0: (v5 * v5), F1: 128, F2: (v5 / ((9223372036854775807) | 1))}
        var v28 uint8 = 220
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    v5 = 42
    var v29 int16 = -17
    v30 := GenS1{F0: v5, F1: -748, F2: v5}
    var v31 int = v5
    if (0.1 <= 0.0) {
        if true {
            v31 += v5
        } else {
            _ = len(v4)
            v32 := GenS4{F0: 907}
            v3 += (v3 * (v29 * v29))
            _ = v32
        }
        var v33 uint16 = 889
        _ = v33
    }
    if true {
        for fi34 := 0; fi34 < 19; fi34++ {
            if true {
                v31 += (v31 ^ v5)
                for fi35 := 0; fi35 < 1; fi35++ {
                    v5 += (-(fi35))
                    v36 := GenS1{F0: (fi34 >> ((2) & 63)), F1: fi34, F2: 42}
                    var v37 int64 = 42
                    switch ("foo bar" + ((("\n\t" + ((("hello" + "foo bar") + ((("\n\t" + "hello") + "hello") + "foo bar")) + "foo bar")) + "\n\t") + "\n\t")) {
                    case "":
                        if false {
                            v5 = 42
                        }
                    case "foo bar":
                        v38 := GenS5{F0: (fi35 & fi34), F1: 2147483648}
                        _ = v38
                    }
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            } else {
                var v39 int = (fi34 / (((fi34 % ((fi34) | 1))) | 1))
                _ = v39
            }
            for fi40 := 0; fi40 < 20; fi40++ {
                switch (v31 + (fi40 * 779)) {
                case 1:
                    switch "a" {
                    case "a":
                        v29 = (^(v3))
                    default:
                        var v41 uint8 = 2
                        _ = v41
                    }
                case 854:
                    if false {
                        var v42 bool = true
                        if (4294967295 < 2147483648) {
                            fmt.Println(v31, fi40)
                            var v43 bool = ((!(3.14159 == (1e10 / 3.14159))) && (!v42))
                            _ = v43
                        } else {
                            v44 := GenS4{F0: 412}
                            _ = v44
                        }
                        _ = v42
                    } else {
                        switch v5 {
                        case 4294967295:
                            v45 := GenS6{F0: "a"}
                            _ = v45
                        case 128:
                            delete(v4, (v31 & (v5 % (((^(v5))) | 1))))
                        case 111:
                            v46 := GenS7{F0: (!true), F1: (v29 / (((v29 | (-(v29)))) | 1)), F2: 0}
                            _ = v46
                        }
                        _ = len(v4)
                    }
                    v31 -= (v5 - (v5 | (-(fi40))))
                case 32768:
                    switch "hello" {
                    case "a":
                        v1 = append(v1, 0)
                    case "foo bar":
                        if true {
                            v5 += 32768
                        } else {
                            v47 := GenS5{F0: 42, F1: 688}
                            var v48 bool = true
                            v49 := GenS7{F0: true, F1: (v3 << ((42) & 63)), F2: 0}
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                    case "":
                        if len(v2) > 0 {
                            v2[0] = ("foo bar" + (("hello" + "foo bar") + ((("hello" + "") + "\n\t") + "hello")))
                        }
                    }
                }
                v31 = fi34
                _ = fi40
            }
            for fi50 := 0; fi50 < 18; fi50++ {
                var v51 int64 = -1
                _ = fi50
                _ = v51
            }
            _ = fi34
        }
        for fi52 := 0; fi52 < 15; fi52++ {
            v53 := GenS7{F0: true, F1: v3, F2: 0}
            _ = fi52
            _ = v53
        }
    }
    var v54 float64 = (1.5 / ((-2.25 - 0.0) - ((1.5 * (((-2.25 / float64(1.5)) / 1e10) + ((0.1 / 1.5) / 3.14159))) - (float64(0.0) - (3.14159 * 3.14159)))))
    v54 = (v54 / 1.0)
    v54 += (v54 * v54)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v54
}
