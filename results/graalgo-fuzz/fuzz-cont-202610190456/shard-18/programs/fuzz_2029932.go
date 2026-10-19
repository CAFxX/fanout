package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 int32; F2 uint }

func main() {
    v1 := []string{"foo bar", "foo bar", "\n\t"}
    var v2 string = ("hello" + ("foo bar" + "\n\t"))
    var v3 uint64 = 1
    v4 := GenS1{F0: 1}
    if true {
        fmt.Println(v2, v2)
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(fi5, v3, v2)
            _ = fi5
        }
    }
    var v6 string = "hello"
    v7 := GenS1{F0: 9223372036854775807}
    switch ((v6 + "foo bar") + (("hello" + v2) + "foo bar")) {
    case "":
        switch v2 {
        case "":
            if false {
                fmt.Println(v6, v3, v3)
                var v8 int8 = -1
                if (true && false) {
                    var v9 uint64 = v3
                    v1 = append(v1, "a")
                    _ = v9
                } else {
                    v10 := GenS2{F0: (v6 + (((v6 + v2) + ("\n\t" + "foo bar")) + v2))}
                    _ = v10
                }
                _ = v8
            }
        }
    case "foo bar":
        var v11 int32 = 1
        _ = v11
    case "\n\t":
        v12 := GenS2{F0: "\n\t"}
        _ = v12
    default:
        var v13 int8 = int8(v3)
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            switch fi14 {
            case 761:
                var v16 uint8 = 0
                _ = v16
            default:
                v2 = v2
            }
            _ = fi15
        }
        v1 = append(v1, v2)
        _ = fi14
    }
    var v17 int16 = 42
    v17 = (v17 >> ((1) & 63))
    var v18 int64 = 555
    var v19 bool = (!true)
    v2 = (v2 + (v6 + v6))
    var v20 uint64 = (v3 / ((0) | 1))
    _ = len(v1)
    if v19 {
        v3 -= v20
    }
    if v19 {
        switch 127 {
        case -2147483648:
            v18 += (v18 << ((1) & 63))
        case 127:
            var v21 float32 = 0.1
            _ = v21
        case -1:
            if ((-(v17)) <= -199) {
                for fi22 := 0; fi22 < 1; fi22++ {
                    v1 = append(v1, "hello")
                    switch (fi22 / (((fi22 ^ (fi22 & (fi22 << ((1) & 63))))) | 1)) {
                    case 887:
                        for fi23 := 0; fi23 < 11; fi23++ {
                            if v19 {
                                v24 := GenS1{F0: v18}
                                fmt.Println(v20, fi23, v2)
                                _ = v24
                            } else {
                                fmt.Println(v17)
                                switch "foo bar" {
                                case "foo bar":
                                    var v25 uint64 = v3
                                    fmt.Println(v20, v19, v6)
                                    var v26 float32 = 1e10
                                    _ = v25
                                    _ = v26
                                case "\n\t":
                                    v2 = (v6 + v2)
                                default:
                                    if false {
                                        var v27 uint8 = 0
                                        v28 := GenS3{F0: true}
                                        var v29 float64 = -2.25
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    } else {
                                        if false {
                                            v30 := GenS3{F0: ((!((-20 == -73) || v19)) || (0.1 != -2.25))}
                                            if ((v18 >> ((504) & 63)) == v18) {
                                                if (2 != 255) {
                                                    fmt.Println(v17, v19)
                                                } else {
                                                    v31 := GenS2{F0: ((("\n\t" + (((("foo bar" + "\n\t") + ("hello" + v6)) + v6) + v2)) + ((v6 + "") + v2)) + "hello")}
                                                    _ = v31
                                                }
                                                var v32 uint = 255
                                                _ = v32
                                            }
                                            _ = v30
                                        } else {
                                            var v33 bool = v19
                                            switch (^(fi23)) {
                                            case 4294967295:
                                                if ((751 >= 32767) && v33) {
                                                    if ((v33 && true) && false) {
                                                        switch (v2 + ((v6 + (v6 + v2)) + v2)) {
                                                        case "a":
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                var v34 uint32 = uint32(v3)
                                                                _ = v34
                                                            }
                                                        case "\n\t":
                                                            if v19 {
                                                                var v35 int64 = v18
                                                                _ = v35
                                                            }
                                                            if (v6 > (v2 + v6)) {
                                                                if (!true) {
                                                                    v18 = (v18 >> ((148) & 63))
                                                                    var v36 int64 = v18
                                                                    _ = v36
                                                                } else {
                                                                    _ = len(v1)
                                                                    var v37 float64 = -2.25
                                                                    _ = v37
                                                                }
                                                                var v38 uint = 255
                                                                v33 = true
                                                                _ = v38
                                                            }
                                                            v39 := GenS3{F0: v33}
                                                            _ = v39
                                                        case "hello":
                                                            _ = len(v1)
                                                        }
                                                    }
                                                    var v40 uint16 = 540
                                                    _ = v40
                                                }
                                            case -9223372036854775808:
                                                if v33 {
                                                    var v41 float32 = float32(((3.14159 - float64(1.5)) - 0.0))
                                                    _ = v41
                                                } else {
                                                    v2 = (((("hello" + "") + v6) + v2) + (v6 + v2))
                                                    v17 = (v17 / ((0) | 1))
                                                }
                                            case 32768:
                                                v42 := GenS2{F0: v2}
                                                _ = v42
                                            }
                                            v43 := GenS2{F0: ("a" + (v6 + v6))}
                                            if false {
                                                var v44 int16 = 1
                                                switch (fi22 - fi23) {
                                                case 1:
                                                    var v45 string = v2
                                                    _ = v45
                                                case 42:
                                                    v46 := GenS3{F0: (!(float64(1.5) <= (0.1 * ((1.5 * (1.5 - 3.14159)) * 1e10))))}
                                                    _ = v46
                                                }
                                                _ = v44
                                            }
                                            _ = v33
                                            _ = v43
                                        }
                                        if false {
                                            fmt.Println(fi22)
                                            var v47 uint64 = (^(v20))
                                            _ = v47
                                        } else {
                                            var v48 bool = (false || v19)
                                            var v49 string = "hello"
                                            if len(v1) > 0 {
                                                v1[0] = (v6 + v6)
                                            }
                                            _ = v48
                                            _ = v49
                                        }
                                    }
                                }
                            }
                            if (!v19) {
                                if len(v1) > 0 {
                                    v1[0] = ("foo bar" + "")
                                }
                                switch fi23 {
                                case 127:
                                    var v50 int32 = 508
                                    _ = v50
                                case 1:
                                    v17 = (-(v17))
                                case 32768:
                                    var v51 uint16 = 42
                                    _ = v51
                                default:
                                    if ((v3 >> ((255) & 63)) <= 42) {
                                        fmt.Println(v6)
                                    }
                                }
                                var v52 int8 = 0
                                _ = v52
                            }
                            if (42 <= 1) {
                                var v53 int = (fi23 & int(v17))
                                var v54 float32 = 3.14159
                                v53 += fi23
                                _ = v53
                                _ = v54
                            } else {
                                v55 := GenS4{F0: -2147483648, F1: 2147483647, F2: 255}
                                _ = v55
                            }
                            _ = fi23
                        }
                    }
                    switch 119 {
                    case 127:
                        _ = len(v1)
                    case -600:
                        for fi56 := 0; fi56 < 20; fi56++ {
                            var v57 uint16 = 32767
                            v58 := GenS1{F0: (^(v18))}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi56
                            _ = v57
                            _ = v58
                        }
                    }
                    _ = fi22
                }
                v2 = ("\n\t" + v2)
            }
        default:
            if (!v19) {
                var v59 int16 = (v17 >> ((211) & 63))
                var v60 int64 = -1
                _ = v59
                _ = v60
            }
        }
        v18 += v18
        if v19 {
            v18 = -9223372036854775808
            for fi61 := 0; fi61 < 13; fi61++ {
                v19 = v19
                _ = fi61
            }
            var v62 uint32 = 65535
            _ = v62
        } else {
            if v19 {
                var v63 float64 = -2.25
                _ = v63
            } else {
                v18 += -561
                var v64 uint64 = (^(v20))
                _ = v64
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
