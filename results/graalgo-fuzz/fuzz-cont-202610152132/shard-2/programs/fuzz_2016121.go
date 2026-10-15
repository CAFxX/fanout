package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 float64 = 1e10
    var v2 uint = 4294967295
    var v3 bool = false
    v4 := make(map[string]int)
    v5 := []uint32{659, 42, 920}
    v2 = v2
    var v6 int8 = 1
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = v3
        _ = fi7
        _ = v8
    }
    var v9 string = ""
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int8 = v6
        _ = fi10
        _ = v11
    }
    if len(v5) > 0 {
        v5[0] = 2
    }
    fmt.Println(v6, v9, v1)
    for fi12 := 0; fi12 < 8; fi12++ {
        if v3 {
            var v13 int32 = 42
            switch "\n\t" {
            case "hello":
                if v3 {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v15 := GenS1{F0: 0}
                        fmt.Println(v6, v9)
                        _ = fi14
                        _ = v15
                    }
                    v16 := GenS1{F0: 1}
                    switch v9 {
                    case "a":
                        v9 = "a"
                    case "\n\t":
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v2 = (^(v2))
                            v18 := GenS1{F0: 0}
                            if true {
                                v19 := GenS1{F0: -886}
                                var v20 int8 = (v6 + 0)
                                _ = v19
                                _ = v20
                            }
                            _ = fi17
                            _ = v18
                        }
                    default:
                        var v21 uint16 = 65535
                        _ = v21
                    }
                    _ = v16
                } else {
                    v22 := GenS1{F0: 671}
                    v23 := GenS1{F0: 1000}
                    _ = v22
                    _ = v23
                }
            case "\n\t":
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 string = ("hello" + "foo bar")
                    if v3 {
                        var v26 int = (fi12 >> (((^(v2))) & 63))
                        _ = v26
                    } else {
                        v2 = v2
                        var v27 uint16 = 528
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
            case "":
                v28 := GenS1{F0: -855}
                _ = v28
            default:
                if (-1 > -1) {
                    switch v9 {
                    case "":
                        v2 = (v2 / (((v2 ^ (v2 + (v2 & v2)))) | 1))
                    default:
                        v6 = (v6 >> ((65535) & 63))
                    }
                }
            }
            _ = len(v5)
            _ = v13
        } else {
            var v29 float64 = v1
            _ = v29
        }
        _ = fi12
    }
    if v3 {
        fmt.Println(v3)
    } else {
        switch (v9 + v9) {
        case "foo bar":
            if (v1 <= v1) {
                v1 = (v1 / (v1 + ((1.5 + (v1 * (v1 + v1))) - float64(-2.25))))
                for fi30 := 0; fi30 < 18; fi30++ {
                    v31 := GenS2{F0: 2}
                    var v32 int64 = -9223372036854775808
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = len(v5)
            } else {
                delete(v4, v9)
                if v3 {
                    if (v3 || false) {
                        fmt.Println(v2, v6, v9)
                    } else {
                        if ((-2.25 - 0.0) >= 0.1) {
                            var v33 string = ("foo bar" + v9)
                            _ = v33
                        } else {
                            switch v9 {
                            case "":
                                var v34 bool = (0 < 1)
                                _ = v34
                            case "foo bar":
                                var v35 string = (v9 + (v9 + v9))
                                _ = v35
                            }
                            var v36 int = 32768
                            _ = v36
                        }
                    }
                }
            }
        }
        for fi37 := 0; fi37 < 3; fi37++ {
            v1 = (3.14159 - (v1 - (((v1 - v1) * v1) - v1)))
            switch (-(fi37)) {
            case -952:
                if ((!(1 < 9223372036854775807)) && ((fi37 << (((v2 - v2)) & 63)) > (fi37 - fi37))) {
                    v38 := GenS3{F0: 2}
                    _ = v38
                }
            }
            switch (^(fi37)) {
            case 32768:
                if (-995 > 767) {
                    for fi39 := 0; fi39 < 14; fi39++ {
                        v2 += 65535
                        _ = fi39
                    }
                }
            case 9223372036854775807:
                var v40 uint16 = 65535
                _ = v40
            case 127:
                _ = len(v4)
            }
            _ = fi37
        }
    }
    v3 = (782 > 0)
    var v41 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v9
}
