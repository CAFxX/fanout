package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint8; F1 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 int16 = 355
    v4 := make(map[int]string)
    if len(v2) > 0 {
        v2[0] = 42
    }
    v5 := GenS1{F0: 65535}
    var v6 int32 = -1
    var v7 uint32 = 4294967295
    if false {
        v3 += 0
        _ = len(v4)
        v8 := GenS2{F0: 42, F1: -534}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch "" {
            case "foo bar":
                if (("hello" + "hello") > "hello") {
                    var v10 int64 = 1
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v7 += (v7 - (v7 << ((65535) & 63)))
                        switch fi9 {
                        case 566:
                            var v12 uint = 2
                            _ = v12
                        }
                        var v13 uint32 = (v7 << ((2) & 63))
                        _ = fi11
                        _ = v13
                    }
                    var v14 uint64 = 42
                    _ = v10
                    _ = v14
                } else {
                    v3 = (v3 << ((0) & 63))
                    v6 = v6
                    _ = len(v2)
                }
            }
            v15 := GenS2{F0: 2, F1: -9223372036854775808}
            if false {
                if (1.5 >= (((((0.1 * -2.25) + 0.0) * 1e10) * 3.14159) + 0.0)) {
                    v7 += (v7 + v7)
                }
            } else {
                var v16 bool = false
                if false {
                    v2 = append(v2, fi9)
                    v7 += v7
                    switch (fi9 - fi9) {
                    case 574:
                        var v17 bool = v16
                        _ = v17
                    }
                } else {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        if false {
                            v6 += v6
                        } else {
                            v19 := GenS2{F0: 920, F1: 1}
                            var v20 int16 = v3
                            var v21 int = fi9
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                        if (108 < 127) {
                            switch ("\n\t" + "") {
                            case "foo bar":
                                var v22 uint = 42
                                _ = v22
                            }
                        } else {
                            v23 := v4[-1]
                            var v24 string = v23
                            var v25 bool = (-128 <= 89)
                            v26 := GenS1{F0: 933}
                            _ = v23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        _ = fi18
                    }
                }
                _ = v16
            }
            _ = fi9
            _ = v15
        }
        var v27 uint = 255
        for fi28 := 0; fi28 < 5; fi28++ {
            if false {
                if false {
                    fmt.Println(fi28, v7)
                    v29 := GenS2{F0: 42, F1: -1}
                    _ = v29
                }
                for fi30 := 0; fi30 < 18; fi30++ {
                    v3 -= v3
                    switch (fi30 >> (((v27 % (((v27 + (v27 - (v27 >> (((v27 << ((v27) & 63))) & 63))))) | 1))) & 63)) {
                    case 970:
                        var v31 float32 = 3.14159
                        _ = v31
                    case 264:
                        if (int32(v7) >= int32(v3)) {
                            v2 = append(v2, fi30)
                            _ = len(v2)
                        } else {
                            v32 := GenS1{F0: 1}
                            v27 = 255
                            fmt.Println(fi28, v27)
                            _ = v32
                        }
                    case 391:
                        v6 = (v6 | v6)
                        if true {
                            v33 := v1["foo bar"]
                            fmt.Println(v27, v27, v33)
                            _ = v33
                        }
                    }
                    v6 = v6
                    _ = fi30
                }
                if (false && false) {
                    v7 = v7
                } else {
                    var v34 float32 = -2.25
                    _ = v34
                }
            }
            _ = fi28
        }
        _ = v27
    }
    fmt.Println(v7, v6)
    var v35 bool = (1.5 > (3.14159 * (((3.14159 + 1e10) + 0.1) / 1e10)))
    var v36 int = 88
    v3 += v3
    var v37 int32 = (v6 / ((958) | 1))
    fmt.Println(v36, v35, v35)
    for fi38 := 0; fi38 < 7; fi38++ {
        if v35 {
            v39 := GenS3{F0: ((("hello" + "") + "hello") + ("" + "a")), F1: 0}
            _ = v39
        }
        var v40 float32 = 1e10
        _ = fi38
        _ = v40
    }
    var v41 bool = (v35 || v35)
    if (!(v7 >= 4294967295)) {
        for fi42 := 0; fi42 < 13; fi42++ {
            for fi43 := 0; fi43 < 6; fi43++ {
                v37 -= (v37 << ((157) & 63))
                var v44 int16 = v3
                _ = fi43
                _ = v44
            }
            _ = fi42
        }
        v45 := GenS4{F0: 1, F1: v37}
        switch (v36 ^ 4294967295) {
        case 128:
            v2 = append(v2, -1)
        case 9223372036854775807:
            _ = len(v4)
        case 191:
            v37 -= v37
        }
        _ = v45
    } else {
        if (v37 == v37) {
            v46 := GenS4{F0: 255, F1: (v37 | (v37 % ((-2147483648) | 1)))}
            v6 = (v37 * -2147483648)
            var v47 float64 = (1.5 - ((((1.5 * 0.0) * ((-2.25 + 0.1) + 1e10)) * float64(0.0)) + (0.1 * 1.5)))
            _ = v46
            _ = v47
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
