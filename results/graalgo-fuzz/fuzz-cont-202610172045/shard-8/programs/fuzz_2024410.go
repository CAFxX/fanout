package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 int64 }

func main() {
    v1 := []int{}
    var v2 int64 = 0
    v3 := make(map[string]string)
    var v4 uint8 = 0
    var v5 uint = 255
    switch (("\n\t" + "hello") + "a") {
    case "\n\t":
        switch "" {
        case "":
            v6 := v3["hello"]
            _ = v6
        case "\n\t":
            var v7 uint8 = (^(v4))
            _ = v7
        case "a":
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(v2)
                _ = fi8
            }
        default:
            fmt.Println(v2, v4)
        }
    }
    var v9 float32 = 3.14159
    v10 := GenS1{F0: -169}
    var v11 bool = true
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v11 {
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: -2147483648}
            _ = fi12
            _ = v13
        }
        v4 = v4
    } else {
        v14 := v3["\n\t"]
        _ = v14
    }
    v5 += v5
    v15 := GenS1{F0: -359}
    v16 := v3[(("" + (("foo bar" + "hello") + "")) + "")]
    var v17 int32 = 42
    if (-1 != int8(v17)) {
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 uint16 = uint16(v4)
            for fi20 := 0; fi20 < 11; fi20++ {
                v2 = (v2 >> (((v5 + (v5 | v5))) & 63))
                var v21 bool = v11
                if (1 < -1) {
                    fmt.Println(v16, v19)
                } else {
                    v4 = (v4 - v4)
                    var v22 uint64 = 42
                    var v23 int64 = (v2 - v2)
                    _ = v22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            if false {
                v24 := GenS1{F0: -2147483648}
                _ = v24
            } else {
                fmt.Println(v17)
                if (v5 < v5) {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        var v26 float32 = v9
                        _ = fi25
                        _ = v26
                    }
                    var v27 int = int(v4)
                    v28 := GenS1{F0: (v17 % ((v17) | 1))}
                    _ = v27
                    _ = v28
                }
            }
            _ = fi18
            _ = v19
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            if false {
                var v30 int16 = -32768
                _ = v30
            } else {
                v4 -= 255
                v31 := GenS1{F0: (v17 >> ((v5) & 63))}
                v5 += (v5 << (((v5 >> ((v5) & 63))) & 63))
                _ = v31
            }
            _ = fi29
        }
    } else {
        var v32 uint32 = 4294967295
        _ = v32
    }
    fmt.Println(v17, v16, v11)
    switch 758 {
    case 127:
        v33 := GenS2{F0: -9223372036854775808, F1: 65535, F2: v2}
        _ = v33
    case -9223372036854775808:
        switch 127 {
        case 32:
            v16 = ((v16 + v16) + v16)
        default:
            if v11 {
                v2 = (v2 ^ v2)
                for fi34 := 0; fi34 < 12; fi34++ {
                    delete(v3, ((v16 + v16) + (v16 + v16)))
                    _ = fi34
                }
            }
        }
    case 160:
        v35 := GenS1{F0: (v17 | v17)}
        _ = v35
    }
    if v11 {
        v11 = (!(-1 == -966))
        v4 = v4
        var v36 int64 = v2
        _ = v36
    } else {
        if ((v5 * v5) >= (v5 - (v5 * (v5 & (^(v5)))))) {
            switch -9223372036854775808 {
            case 2147483647:
                v9 = -2.25
            case 4294967295:
                var v37 int = -2147483648
                _ = v37
            case 9223372036854775807:
                v1 = append(v1, 979)
            }
        } else {
            var v38 uint = 2
            v17 = int32(v2)
            _ = v38
        }
        var v39 int8 = -128
        v2 -= 9223372036854775807
        _ = v39
    }
    var v40 float32 = 0.1
    v5 = 689
    switch v16 {
    case "a":
        fmt.Println(v16, v40, v40)
    }
    for fi41 := 0; fi41 < 20; fi41++ {
        for fi42 := 0; fi42 < 12; fi42++ {
            var v43 uint = v5
            var v44 string = ("hello" + ((("" + "a") + v16) + "a"))
            if len(v1) > 0 {
                v1[0] = fi42
            }
            switch (fi42 + fi41) {
            case 2147483647:
                v45 := GenS2{F0: int(v43), F1: 42, F2: -1}
                _ = v45
            default:
                delete(v3, "a")
            }
            v5 -= v43
            _ = fi42
            _ = v43
            _ = v44
        }
        for fi46 := 0; fi46 < 6; fi46++ {
            v47 := v3[(((v16 + v16) + v16) + v16)]
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v48 := GenS1{F0: 0}
            _ = fi46
            _ = v47
            _ = v48
        }
        for fi49 := 0; fi49 < 5; fi49++ {
            var v50 string = "a"
            v51 := GenS1{F0: (v17 ^ (v17 ^ v17))}
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
