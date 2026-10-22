package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]int)
    var v3 uint64 = 117
    v4 := make(map[string]uint32)
    switch "" {
    case "":
        v1 += (v1 << ((307) & 63))
    default:
        var v5 int32 = -1
        _ = v5
    }
    fmt.Println(v1, v1)
    if true {
        v6 := v4["a"]
        for fi7 := 0; fi7 < 4; fi7++ {
            v6 -= 159
            switch fi7 {
            case -372:
                var v8 bool = (65535 < 255)
                _ = v8
            case 0:
                _ = len(v2)
            case 103:
                v6 += 0
            default:
                _ = len(v2)
            }
            if (!true) {
                var v9 uint16 = 0
                fmt.Println(fi7)
                _ = v9
            } else {
                for fi10 := 0; fi10 < 15; fi10++ {
                    if ((0.0 <= 0.0) && true) {
                        fmt.Println(fi7, v3)
                    }
                    _ = fi10
                }
                fmt.Println(v1, fi7, v6)
            }
            _ = fi7
        }
        _ = v6
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint64 = v3
                _ = fi12
                _ = v13
            }
            var v14 int64 = -217
            _ = fi11
            _ = v14
        }
        fmt.Println(v1)
        v3 = 196
    }
    v1 = (v1 | v1)
    var v15 bool = (!(((((3.14159 + 1e10) * 1e10) + 1e10) - 1e10) != float64((1.5 / 1e10))))
    v1 += v1
    v3 = (^(v3))
    for fi16 := 0; fi16 < 4; fi16++ {
        switch fi16 {
        case 156:
            if ((v1 <= -832) || (137 > -179)) {
                v17 := GenS1{F0: 65535}
                if v15 {
                    v3 = v3
                    _ = len(v4)
                }
                _ = v17
            } else {
                if ((2 > 1) && v15) {
                    _ = len(v4)
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 uint32 = 4294967295
                        switch ((("" + "a") + ("\n\t" + "hello")) + "hello") {
                        case "hello":
                            if (v15 || (3.14159 != 0.1)) {
                                var v20 bool = (!(812 > 4294967295))
                                _ = v20
                            }
                        }
                        fmt.Println(v19, v1, fi18)
                        _ = fi18
                        _ = v19
                    }
                }
                var v21 uint = 424
                v15 = ((^(v3)) >= uint64(v21))
                _ = v21
            }
        case 32768:
            var v22 bool = (v1 > (v1 + v1))
            _ = v22
        }
        if false {
            v1 += 42
            fmt.Println(v1)
            v23 := []int{0, (fi16 << ((110) & 63))}
            _ = v23
        } else {
            switch (-(fi16)) {
            case -9223372036854775808:
                if false {
                    var v24 uint16 = 1
                    v25 := GenS1{F0: 42}
                    _ = v24
                    _ = v25
                } else {
                    if v15 {
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v15, v15, v15)
                        case "a":
                            v26 := GenS1{F0: 255}
                            _ = v26
                        case "foo bar":
                            var v27 uint32 = 4294967295
                            _ = v27
                        }
                        _ = len(v4)
                    } else {
                        var v28 float32 = (float32(1e10) * 0.0)
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 3; fi29++ {
                        fmt.Println(v3, v3)
                        _ = fi29
                    }
                }
            case 9223372036854775807:
                v4[(("\n\t" + "hello") + (("" + ("\n\t" + "hello")) + ("hello" + "a")))] = 42
            case 32768:
                v30 := GenS1{F0: 32767}
                _ = v30
            default:
                switch "" {
                case "hello":
                    switch (fi16 & 0) {
                    case -824:
                        v3 = 4294967295
                    case 603:
                        var v31 int64 = -9223372036854775808
                        _ = v31
                    }
                default:
                    switch fi16 {
                    case 279:
                        for fi32 := 0; fi32 < 7; fi32++ {
                            var v33 uint64 = (v3 / ((955) | 1))
                            _ = fi32
                            _ = v33
                        }
                    case 667:
                        v15 = true
                    case 128:
                        v34 := []int{}
                        _ = v34
                    default:
                        v1 = -37
                    }
                }
            }
        }
        _ = fi16
    }
    v35 := v2[2147483647]
    var v36 uint32 = 42
    var v37 string = "hello"
    var v38 int32 = 397
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
