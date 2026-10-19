package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    var v1 float32 = 0.0
    var v2 int = 127
    v3 := []string{"\n\t", ("\n\t" + "a"), ""}
    var v4 string = ("" + "foo bar")
    var v5 int64 = 893
    if ((v5 ^ 9223372036854775807) <= v5) {
        var v6 int8 = 100
        _ = v6
    } else {
        var v7 int64 = v5
        v4 = v4
        _ = v7
    }
    switch v4 {
    case "":
        v8 := GenS1{F0: 1, F1: -32768, F2: v5}
        _ = v8
    }
    var v9 string = v4
    v3 = append(v3, "a")
    for fi10 := 0; fi10 < 11; fi10++ {
        v2 += v2
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: 4294967295, F1: -129, F2: v5}
            var v13 uint32 = 65535
            v4 = "hello"
            fmt.Println(fi10)
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := GenS2{F0: 343}
        _ = fi10
        _ = v14
    }
    var v15 uint32 = 4294967295
    for fi16 := 0; fi16 < 3; fi16++ {
        if false {
            switch v9 {
            case "hello":
                if len(v3) > 0 {
                    v3[0] = v9
                }
            case "\n\t":
                v1 += v1
            }
        }
        _ = fi16
    }
    v3 = append(v3, (v9 + v4))
    for fi17 := 0; fi17 < 2; fi17++ {
        switch fi17 {
        case 1:
            for fi18 := 0; fi18 < 1; fi18++ {
                _ = len(v3)
                switch v9 {
                case "\n\t":
                    var v19 int32 = 42
                    _ = v19
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "a":
                    switch (v4 + v9) {
                    case "foo bar":
                        var v20 float64 = -2.25
                        _ = v20
                    case "a":
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v15 = v15
                        }
                    }
                }
                v3 = append(v3, v9)
                _ = fi18
            }
        case 9223372036854775807:
            v21 := GenS1{F0: 0, F1: 42, F2: v5}
            _ = v21
        }
        _ = len(v3)
        _ = fi17
    }
    switch v9 {
    case "foo bar":
        switch -1 {
        case -2147483648:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            switch v2 {
            case 76:
                v1 -= v1
            default:
                switch ("hello" + "") {
                case "":
                    if false {
                        switch v2 {
                        case 32768:
                            _ = len(v3)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    } else {
                        if (int16(v5) < -939) {
                            v22 := GenS2{F0: 908}
                            if (42 == 2) {
                                v4 = "foo bar"
                                v5 = v5
                                v4 = ((v9 + v4) + v9)
                            }
                            v3 = append(v3, (v4 + (v4 + v4)))
                            _ = v22
                        } else {
                            var v23 uint64 = 0
                            _ = v23
                        }
                        v24 := GenS1{F0: 707, F1: 42, F2: (-(v5))}
                        _ = v24
                    }
                }
            }
        }
    }
    if ((42 <= v2) || true) {
        v25 := GenS3{F0: 42, F1: ((2147483647 > 0) || false)}
        switch v2 {
        case 127:
            switch v9 {
            case "":
                var v26 string = v4
                _ = v26
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v27 := GenS1{F0: 1, F1: 0, F2: 42}
        _ = v25
        _ = v27
    }
    if true {
        v28 := GenS1{F0: 42, F1: 1, F2: (v5 + v5)}
        _ = v28
    } else {
        if len(v3) > 0 {
            v3[0] = ""
        }
        fmt.Println(v4, v2)
    }
    v5 = int64(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
