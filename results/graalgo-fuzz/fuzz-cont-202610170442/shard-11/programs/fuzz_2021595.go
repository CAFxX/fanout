package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 274
    v3 := make(map[string]int)
    v4 := []int{42, 0}
    var v5 uint64 = 825
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 804, F1: -32768, F2: 343}
        _ = len(v1)
        v5 += (v5 + (v5 >> ((42) & 63)))
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: v5, F1: 1}
    switch "foo bar" {
    case "a":
        var v10 float64 = 1.5
        _ = v10
    case "\n\t":
        v11 := v6[533]
        _ = v11
    default:
        switch 842 {
        case 4294967295:
            v2 -= v2
        case 42:
            v12 := GenS3{F0: 1, F1: 102, F2: 32767}
            _ = v12
        case -639:
            _ = len(v4)
        default:
            var v13 uint16 = 1
            _ = v13
        }
    }
    var v14 uint16 = 55
    if (-2.25 < 1e10) {
        v15 := GenS2{F0: v5, F1: 981}
        if (-1 > -1) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                switch ("a" + "hello") {
                case "foo bar":
                    var v16 int32 = -850
                    _ = v16
                case "":
                    v17 := v1[1]
                    _ = v17
                }
                v18 := GenS2{F0: (^(v5)), F1: 976}
                switch (("foo bar" + (((("" + "a") + (("a" + "") + (("foo bar" + ("\n\t" + "foo bar")) + "a"))) + "\n\t") + "\n\t")) + ("foo bar" + "")) {
                case "\n\t":
                    fmt.Println(v2, v2, v2)
                case "hello":
                    v1[-870] = 1
                case "foo bar":
                    v2 = v2
                }
                _ = v18
            } else {
                _ = len(v4)
                switch 2147483647 {
                case 127:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 1:
                    v14 -= 255
                default:
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                }
                v14 -= 65535
            }
            fmt.Println(v2)
        }
        fmt.Println(v2, v5)
        _ = v15
    } else {
        var v19 float64 = (1.5 * 1e10)
        v14 -= v14
        for fi20 := 0; fi20 < 8; fi20++ {
            if (false || true) {
                v5 -= v5
                _ = len(v4)
                var v21 uint16 = 836
                _ = v21
            }
            var v22 uint8 = 42
            var v23 uint32 = 2
            _ = fi20
            _ = v22
            _ = v23
        }
        _ = v19
    }
    _ = len(v1)
    v4 = append(v4, 1)
    switch 40 {
    case 1:
        var v24 float64 = -2.25
        _ = v24
    case 9223372036854775807:
        v14 -= v14
    }
    switch "\n\t" {
    case "hello":
        fmt.Println(v5)
    }
    switch -836 {
    case 4294967295:
        var v25 uint16 = 42
        _ = v25
    case 2147483647:
        v26 := GenS2{F0: (v5 + v5), F1: 496}
        _ = v26
    case 9223372036854775807:
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 int64 = 42
            switch fi27 {
            case 9223372036854775807:
                if true {
                    v4 = append(v4, int(v2))
                }
            case 1:
                v29 := GenS4{F0: 270}
                _ = v29
            }
            _ = fi27
            _ = v28
        }
    }
    var v30 float32 = 1.5
    for fi31 := 0; fi31 < 5; fi31++ {
        switch int(v14) {
        case 32768:
            fmt.Println(fi31)
        default:
            _ = len(v1)
        }
        v32 := GenS5{F0: 0, F1: 255, F2: v5}
        if (18446744073709551615 == 1) {
            var v33 int32 = 0
            var v34 float32 = 1e10
            _ = v33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    if true {
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 uint64 = (^(v5))
            v30 = v30
            _ = fi35
            _ = v36
        }
        switch ((("\n\t" + (("foo bar" + "a") + "foo bar")) + ("" + "")) + "") {
        case "a":
            v37 := GenS4{F0: -1}
            _ = v37
        default:
            v38 := v6[127]
            _ = v38
        }
        if true {
            v5 = 4294967295
            v30 -= v30
            switch -686 {
            case -9223372036854775808:
                var v39 uint8 = 1
                _ = v39
            }
        } else {
            fmt.Println(v30, v30, v5)
        }
    } else {
        if true {
            if (true || true) {
                v2 -= 0
                v4 = append(v4, -1)
                v5 = v5
            } else {
                for fi40 := 0; fi40 < 10; fi40++ {
                    v3["hello"] = (fi40 << ((1) & 63))
                    _ = fi40
                }
                v30 = v30
                for fi41 := 0; fi41 < 10; fi41++ {
                    if len(v4) > 0 {
                        v4[0] = fi41
                    }
                    if (!(!false)) {
                        v30 = ((v30 - -2.25) * float32(1.5))
                        v42 := GenS6{F0: 2, F1: -9223372036854775808}
                        if false {
                            _ = len(v3)
                        }
                        _ = v42
                    }
                    _ = fi41
                }
            }
        }
        fmt.Println(v30, v14, v30)
    }
    _ = len(v1)
    var v43 uint32 = 587
    for fi44 := 0; fi44 < 18; fi44++ {
        v45 := v6[(^(fi44))]
        var v46 uint64 = (v5 * (v5 * (v5 >> ((uint(v5)) & 63))))
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v9
}
