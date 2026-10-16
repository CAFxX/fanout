package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint8 = 50
    v2 := []uint32{909}
    var v3 uint = 1
    var v4 uint = v3
    for fi5 := 0; fi5 < 2; fi5++ {
        v4 -= (v3 & v4)
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = 32767
            v8 := GenS1{F0: 1}
            fmt.Println(fi6)
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    v1 -= uint8(v3)
    fmt.Println(v1, v4)
    switch 127 {
    case 2147483647:
        v2 = append(v2, 2147483647)
    case 32767:
        if ((127 == 100) || (-2.25 < (3.14159 * 0.0))) {
            v3 = (^(v3))
            var v9 uint = v3
            if ((!false) || false) {
                var v10 uint64 = 1
                var v11 int32 = 42
                v12 := GenS1{F0: 255}
                _ = v10
                _ = v11
                _ = v12
            } else {
                fmt.Println(v9, v1)
            }
            _ = v9
        }
    case -392:
        var v13 int16 = 0
        _ = v13
    default:
        if ((v3 & (v3 | v4)) >= v4) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v14 int64 = -903
            _ = v14
        }
    }
    if true {
        v1 = v1
        switch 0 {
        case 553:
            if false {
                switch ("hello" + "hello") {
                case "a":
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        var v15 bool = false
                        _ = v15
                    case "":
                        for fi16 := 0; fi16 < 8; fi16++ {
                            var v17 uint32 = 1
                            fmt.Println(v4)
                            v18 := GenS1{F0: 32767}
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    case "a":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "foo bar":
                    var v19 float32 = 3.14159
                    _ = v19
                case "\n\t":
                    for fi20 := 0; fi20 < 13; fi20++ {
                        var v21 float64 = ((0.1 / float64(-2.25)) * 1.5)
                        _ = fi20
                        _ = v21
                    }
                }
                v4 += (v4 & v3)
            }
        default:
            v22 := GenS1{F0: 2}
            _ = v22
        }
        if len(v2) > 0 {
            v2[0] = 65535
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 = v3
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v23 := GenS2{F0: -443}
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 16; fi25++ {
            v3 = v3
            switch fi25 {
            case 4294967295:
                v26 := GenS3{F0: 2}
                _ = v26
            default:
                if false {
                    var v27 float32 = 1.5
                    _ = v27
                } else {
                    switch ("a" + "\n\t") {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        v3 += v4
                    default:
                        v3 -= v4
                    }
                    if (!(1000 <= 85)) {
                        v2 = append(v2, 52)
                    } else {
                        if (!false) {
                            v3 = 585
                        }
                    }
                    var v28 uint32 = 997
                    _ = v28
                }
            }
            switch "" {
            case "foo bar":
                var v29 int16 = 42
                _ = v29
            }
            _ = fi25
        }
        _ = fi24
    }
    var v30 int32 = 42
    v2 = append(v2, 808)
    v4 -= 255
    switch 32768 {
    case 281:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 32767:
        var v31 int32 = (v30 ^ 1)
        _ = v31
    default:
        v32 := GenS2{F0: int32(v1)}
        _ = v32
    }
    v4 = v4
    for fi33 := 0; fi33 < 14; fi33++ {
        switch (("\n\t" + "foo bar") + "hello") {
        case "":
            fmt.Println(fi33, v1, v4)
        default:
            v4 -= 696
        }
        var v34 uint = 395
        _ = fi33
        _ = v34
    }
    var v35 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v35
    _ = v4
}
