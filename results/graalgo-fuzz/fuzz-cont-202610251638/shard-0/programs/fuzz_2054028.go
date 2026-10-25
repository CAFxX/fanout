package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint32; F2 string }
type GenS3 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    v2 := []int{32768, 2147483647, -2147483648, 128}
    v3 := make(map[uint32]int)
    var v4 float32 = 0.1
    var v5 uint32 = 429
    var v6 bool = (868 == 1)
    v4 += v4
    switch (("\n\t" + "foo bar") + "foo bar") {
    case "\n\t":
        switch "" {
        case "foo bar":
            if (32768 == 0) {
                _ = len(v1)
                switch "a" {
                case "\n\t":
                    for fi7 := 0; fi7 < 16; fi7++ {
                        if (!(-2147483648 == 439)) {
                            v8 := GenS1{F0: 42}
                            delete(v3, v5)
                            _ = v8
                        } else {
                            var v9 float32 = v4
                            _ = v9
                        }
                        _ = fi7
                    }
                default:
                    v5 -= (v5 >> ((1) & 63))
                }
            } else {
                for fi10 := 0; fi10 < 8; fi10++ {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        if (-927 >= fi10) {
                            fmt.Println(v5, fi10)
                            if v6 {
                                fmt.Println(fi10)
                            } else {
                                if v6 {
                                    _ = len(v3)
                                } else {
                                    v12 := GenS2{F0: v6, F1: v5, F2: (("foo bar" + "hello") + "\n\t")}
                                    var v13 uint32 = (v5 >> ((0) & 63))
                                    v14 := GenS2{F0: false, F1: v5, F2: ("\n\t" + "foo bar")}
                                    _ = v12
                                    _ = v13
                                    _ = v14
                                }
                                var v15 float32 = (v4 * (v4 * v4))
                                fmt.Println(v5)
                                _ = v15
                            }
                            var v16 float32 = ((v4 + v4) + 0.1)
                            _ = v16
                        } else {
                            v17 := GenS1{F0: 0}
                            v18 := GenS1{F0: -32768}
                            v19 := GenS2{F0: true, F1: v5, F2: "a"}
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi11
                    }
                    _ = fi10
                }
            }
        case "":
            if (42 == 255) {
                v5 -= v5
            }
        default:
            for fi20 := 0; fi20 < 4; fi20++ {
                switch (fi20 << ((255) & 63)) {
                case 42:
                    v6 = false
                }
                _ = fi20
            }
        }
    case "":
        switch 4294967295 {
        case 32768:
            fmt.Println(v4, v5)
        case 128:
            var v21 bool = (!true)
            _ = v21
        }
    case "a":
        for fi22 := 0; fi22 < 1; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := v3[v5]
                _ = fi23
                _ = v24
            }
            switch fi22 {
            case -2147483648:
                switch ("a" + "") {
                case "hello":
                    if true {
                        v25 := GenS2{F0: v6, F1: (^(v5)), F2: "\n\t"}
                        _ = v25
                    } else {
                        delete(v3, 490)
                    }
                case "a":
                    var v26 string = "a"
                    _ = v26
                default:
                    v6 = v6
                }
            case 32767:
                var v27 float64 = 1.5
                _ = v27
            case 9223372036854775807:
                for fi28 := 0; fi28 < 5; fi28++ {
                    v29 := GenS3{F0: 1}
                    _ = fi28
                    _ = v29
                }
            }
            var v30 float64 = float64(v4)
            _ = fi22
            _ = v30
        }
    }
    var v31 int16 = 1
    var v32 uint = 65535
    v5 += v5
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 int16 = (v31 / ((v31) | 1))
        var v35 int64 = 1
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS2{F0: v6, F1: v5, F2: ("foo bar" + "a")}
    switch -2147483648 {
    case 1:
        switch 4294967295 {
        case 127:
            var v37 uint8 = 1
            _ = v37
        case -270:
            fmt.Println(v5)
        default:
            v38 := v3[v5]
            _ = v38
        }
    }
    v3[v5] = int(v32)
    for fi39 := 0; fi39 < 4; fi39++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi39
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v5 -= 65535
        var v41 uint32 = (v5 + 0)
        v42 := GenS3{F0: v31}
        _ = fi40
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 int = 4294967295
        _ = fi43
        _ = v44
    }
    var v45 uint = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
