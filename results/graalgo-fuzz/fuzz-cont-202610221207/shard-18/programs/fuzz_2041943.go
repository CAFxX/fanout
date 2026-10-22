package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []int{795, 128}
    v2 := make(map[string]bool)
    var v3 bool = ((true || (42 != -1)) || false)
    var v4 int16 = 641
    for fi5 := 0; fi5 < 18; fi5++ {
        if v3 {
            fmt.Println(v3, v3, fi5)
            if v3 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi6 := 0; fi6 < 16; fi6++ {
                    var v7 int16 = v4
                    _ = fi6
                    _ = v7
                }
            } else {
                fmt.Println(v3)
            }
            switch (("\n\t" + "a") + "\n\t") {
            case "hello":
                v4 -= v4
            case "a":
                v4 = v4
            case "":
                fmt.Println(fi5, fi5, v3)
            }
        } else {
            v4 += (^(v4))
            for fi8 := 0; fi8 < 9; fi8++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                delete(v2, "\n\t")
                var v9 int32 = 0
                _ = fi8
                _ = v9
            }
        }
        _ = fi5
    }
    if v3 {
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint32 = 2
            if len(v1) > 0 {
                v1[0] = -851
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint = 2
                var v14 float64 = 1.5
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        var v15 int64 = -9223372036854775808
        _ = v15
    }
    var v16 float64 = float64(0.0)
    v17 := GenS1{F0: 32767}
    fmt.Println(v3, v3, v4)
    if (0 == 1) {
        v3 = false
        if (v4 != 32767) {
            v4 = (-(v4))
        }
        switch 2147483647 {
        case 32767:
            for fi18 := 0; fi18 < 19; fi18++ {
                if v3 {
                    fmt.Println(v3, v16, v16)
                    if ((!true) || v3) {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 uint = 42
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(v4, v16, v16)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                v4 -= (v4 - 32767)
                _ = fi18
            }
        case -563:
            var v21 uint = 118
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            switch "a" {
            case "foo bar":
                switch ("a" + "foo bar") {
                case "foo bar":
                    switch (("\n\t" + "hello") + ("foo bar" + "\n\t")) {
                    case "\n\t":
                        v4 = -426
                    }
                case "a":
                    for fi23 := 0; fi23 < 11; fi23++ {
                        if ((2147483647 != 468) || (fi22 == (-(fi23)))) {
                            var v24 uint32 = 218
                            v25 := GenS1{F0: 2}
                            var v26 int16 = v4
                            if (!v3) {
                                switch -1 {
                                case 175:
                                    var v27 uint = 255
                                    _ = v27
                                case 42:
                                    var v28 int = (fi22 >> ((0) & 63))
                                    _ = v28
                                }
                            }
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        if (!((v3 && (!(97 > 255))) && v3)) {
                            var v29 bool = (2 >= 351)
                            v4 += (-(v4))
                            var v30 string = (("foo bar" + "") + "a")
                            _ = v29
                            _ = v30
                        }
                        v4 = (v4 ^ 1)
                        var v31 int8 = 23
                        _ = fi23
                        _ = v31
                    }
                case "\n\t":
                    for fi32 := 0; fi32 < 18; fi32++ {
                        var v33 uint = 296
                        var v34 int8 = 86
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                }
            case "hello":
                for fi35 := 0; fi35 < 16; fi35++ {
                    var v36 float32 = 0.0
                    _ = fi35
                    _ = v36
                }
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                switch fi22 {
                case 1:
                    v16 = (v16 / 1.5)
                }
            }
            _ = fi22
        }
        var v37 int32 = -2147483648
        _ = v37
    }
    _ = len(v1)
    var v38 int32 = -1
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 float64 = -2.25
        v40 += (3.14159 * ((v16 - (0.0 + v16)) * v16))
        v41 := GenS1{F0: 736}
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
