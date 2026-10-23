package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int64 = -665
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{}
    fmt.Println(v1, v2)
    v2 = v2
    fmt.Println(v2, v2)
    var v4 int = int(v1)
    if (-106 <= 127) {
        v1 = 0
        v1 = v1
    } else {
        for fi5 := 0; fi5 < 19; fi5++ {
            for fi6 := 0; fi6 < 20; fi6++ {
                switch "foo bar" {
                case "":
                    v7 := GenS1{F0: 127}
                    _ = v7
                case "\n\t":
                    _ = len(v3)
                default:
                    if false {
                        var v8 int16 = -1
                        _ = v8
                    } else {
                        v2 -= 1
                        v9 := GenS1{F0: 1}
                        _ = v9
                    }
                }
                fmt.Println(fi5, fi5, v4)
                if false {
                    v4 -= (fi6 % ((v4) | 1))
                } else {
                    if true {
                        v10 := GenS2{F0: 42, F1: 42, F2: 1}
                        _ = v10
                    } else {
                        if false {
                            _ = len(v3)
                        }
                        var v11 int8 = 1
                        _ = v11
                    }
                    v1 -= (-(v1))
                }
                _ = fi6
            }
            v12 := GenS2{F0: 4294967295, F1: 128, F2: v1}
            var v13 uint = 1
            _ = fi5
            _ = v12
            _ = v13
        }
        switch ("a" + "hello") {
        case "foo bar":
            switch -9223372036854775808 {
            case 0:
                var v14 uint16 = 0
                _ = v14
            }
        case "\n\t":
            v4 = v4
        case "":
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS1{F0: 2}
                fmt.Println(v2)
                _ = fi15
                _ = v16
            }
        default:
            switch (v4 + v4) {
            case 177:
                if len(v3) > 0 {
                    v3[0] = 2
                }
            case 282:
                var v17 uint32 = 129
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v4 = v4
        _ = fi18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 float64 = 1.5
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 1}
    v22 := GenS1{F0: 42}
    v1 = int64(v4)
    var v23 int64 = v1
    switch "foo bar" {
    case "a":
        var v24 uint = 565
        _ = v24
    case "\n\t":
        v4 = (v4 ^ 128)
    default:
        for fi25 := 0; fi25 < 13; fi25++ {
            if (true && true) {
                v26 := GenS3{F0: 127}
                for fi27 := 0; fi27 < 9; fi27++ {
                    var v28 string = ("" + "\n\t")
                    v29 := GenS2{F0: 615, F1: 127, F2: v1}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = v26
            } else {
                v4 += int(v1)
                v30 := GenS1{F0: 175}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v30
            }
            v31 := GenS2{F0: 42, F1: 127, F2: v23}
            fmt.Println(fi25, v23, v1)
            _ = fi25
            _ = v31
        }
    }
    v1 += (v1 << ((18446744073709551615) & 63))
    fmt.Println(v23)
    if true {
        v4 = (-(v4))
        if false {
            if len(v3) > 0 {
                v3[0] = 963
            }
            fmt.Println(v23)
            for fi32 := 0; fi32 < 8; fi32++ {
                if false {
                    v4 += v4
                } else {
                    var v33 int = (v4 % ((-789) | 1))
                    var v34 float32 = (1.5 * 3.14159)
                    _ = v33
                    _ = v34
                }
                v1 -= (v23 - v23)
                fmt.Println(v2, v2)
                _ = fi32
            }
        } else {
            for fi35 := 0; fi35 < 14; fi35++ {
                fmt.Println(v1, v4)
                fmt.Println(v1)
                _ = fi35
            }
        }
    }
    var v36 float64 = 3.14159
    var v37 int = (v4 % ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
