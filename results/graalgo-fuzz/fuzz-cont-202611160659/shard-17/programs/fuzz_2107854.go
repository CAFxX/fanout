package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint16 }
type GenS2 struct { F0 string; F1 string; F2 int32 }
type GenS3 struct { F0 string }

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 1
    var v3 uint32 = 2
    var v4 uint8 = 255
    v5 := []uint32{(^(v3))}
    var v6 uint32 = (v3 + (v3 + (v3 + v3)))
    fmt.Println(v3, v6, v2)
    if (true || true) {
        if false {
            var v7 uint16 = 2
            _ = v7
        } else {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 uint32 = 2
                v3 -= v9
                var v10 string = "a"
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 int32 = 1
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v2)
                var v13 uint = uint(v2)
                var v14 float32 = (3.14159 * (((3.14159 + (-2.25 / 1e10)) - -2.25) / -2.25))
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            if len(v5) > 0 {
                v5[0] = uint32(v2)
            }
            _ = fi15
        }
        v5 = append(v5, v6)
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v3, v4, v2)
        if (255 == 18446744073709551615) {
            for fi17 := 0; fi17 < 4; fi17++ {
                if true {
                    var v18 float32 = (0.1 * -2.25)
                    _ = v18
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = fi17
            }
            for fi19 := 0; fi19 < 4; fi19++ {
                v3 += 2147483647
                var v20 uint16 = 626
                _ = fi19
                _ = v20
            }
        }
        _ = fi16
    }
    var v21 float64 = (v1 + 1e10)
    for fi22 := 0; fi22 < 18; fi22++ {
        v2 += (v2 / (((v2 - 42)) | 1))
        var v23 uint16 = 2
        if (v23 == v23) {
            v24 := GenS1{F0: 42, F1: "foo bar", F2: 0}
            switch fi22 {
            case 9223372036854775807:
                v23 = (v23 - (v23 - (v23 & v23)))
            case 327:
                var v25 int = -1
                _ = v25
            case -2147483648:
                v23 = uint16(v6)
            }
            var v26 bool = (v21 <= ((v1 * (v1 + (v21 + float64((float32(v21) * 3.14159))))) - v1))
            _ = v24
            _ = v26
        } else {
            v27 := GenS1{F0: 42, F1: "foo bar", F2: (v23 * (^(v23)))}
            v4 -= 128
            _ = v27
        }
        _ = fi22
        _ = v23
    }
    switch 4294967295 {
    case 0:
        v1 += 0.1
    }
    v28 := GenS1{F0: -2147483648, F1: "\n\t", F2: 0}
    var v29 int = 128
    if true {
        switch (^(v29)) {
        case -996:
            var v30 int32 = 0
            _ = v30
        case 4294967295:
            var v31 uint32 = (^(v6))
            _ = v31
        default:
            for fi32 := 0; fi32 < 20; fi32++ {
                fmt.Println(v2, v4)
                var v33 int = (v29 ^ (v29 >> ((65535) & 63)))
                switch ("\n\t" + ("" + "foo bar")) {
                case "\n\t":
                    fmt.Println(v2, v2)
                }
                _ = fi32
                _ = v33
            }
        }
    } else {
        v34 := GenS2{F0: "", F1: ("\n\t" + ("hello" + "a")), F2: 843}
        var v35 float32 = (3.14159 * ((((-2.25 * 1e10) - 1e10) / (-2.25 + (1e10 / (0.0 + (0.1 / -2.25))))) / 1e10))
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 10; fi36++ {
        v3 = v3
        v37 := GenS3{F0: "\n\t"}
        var v38 uint32 = v3
        _ = fi36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        switch v29 {
        case 32768:
            v2 -= 430
        case 4294967295:
            var v40 string = "foo bar"
            _ = v40
        }
        _ = fi39
    }
    v41 := GenS2{F0: "\n\t", F1: "foo bar", F2: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
