package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint64 = 697
    var v2 uint32 = 687
    var v3 int16 = 32767
    var v4 uint = 4294967295
    v5 := make(map[string]bool)
    var v6 float32 = ((1.5 * ((float32(-2.25) - 1.5) - 0.1)) - 0.0)
    for fi7 := 0; fi7 < 9; fi7++ {
        if (false && (0 <= 42)) {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 bool = (-9223372036854775808 >= 42)
                var v10 uint16 = 65535
                _ = fi8
                _ = v10
                _ = v9
            }
            if true {
                v5["\n\t"] = true
            } else {
                var v11 int8 = 71
                _ = v11
            }
        }
        switch "" {
        case "":
            var v12 uint64 = (v1 | 2)
            _ = v12
        default:
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS1{F0: 1}
                _ = fi13
                _ = v14
            }
        }
        _ = fi7
    }
    v15 := GenS1{F0: -128}
    v16 := GenS1{F0: 42}
    fmt.Println(v4, v1)
    if true {
        v1 = (v1 & (v1 << ((v4) & 63)))
    } else {
        fmt.Println(v3)
        v4 -= 827
        var v17 bool = true
        _ = v17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(v4, v6, fi18)
        _ = fi18
    }
    v19 := GenS1{F0: -1}
    if false {
        v20 := GenS2{F0: 99, F1: -1}
        _ = v20
    } else {
        if ((1 < -302) || true) {
            v4 += 1
            switch "foo bar" {
            case "a":
                v21 := GenS2{F0: int(v1), F1: -1}
                _ = v21
            case "\n\t":
                var v22 float64 = 0.1
                _ = v22
            }
            var v23 string = "hello"
            _ = v23
        } else {
            switch "foo bar" {
            case "hello":
                fmt.Println(v3)
            case "\n\t":
                for fi24 := 0; fi24 < 16; fi24++ {
                    switch (-(fi24)) {
                    case 128:
                        if (!true) {
                            v4 = 1
                        }
                    case 785:
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v5["hello"] = false
                            _ = fi25
                        }
                    default:
                        switch ("hello" + "hello") {
                        case "\n\t":
                            var v26 int8 = 100
                            _ = v26
                        case "foo bar":
                            var v27 float32 = v6
                            _ = v27
                        default:
                            v4 = (v4 / ((64) | 1))
                        }
                    }
                    v5["hello"] = (42 > 127)
                    _ = fi24
                }
            }
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            if (!(false || ((false && (1 != 127)) || (2147483647 == -76)))) {
                var v29 uint64 = v1
                v30 := []int{fi28, 32768, (fi28 << ((uint(v2)) & 63)), (^(fi28))}
                _ = v29
                _ = v30
            }
            _ = fi28
        }
        v6 -= 0.1
    }
    v1 -= 284
    for fi31 := 0; fi31 < 2; fi31++ {
        switch fi31 {
        case 1:
            v6 += ((float32(3.14159) / v6) + v6)
        case 4294967295:
            v32 := []int{fi31, (-(fi31)), fi31}
            _ = v32
        case 9223372036854775807:
            var v33 float64 = -2.25
            _ = v33
        default:
            v34 := GenS3{F0: false, F1: (v3 ^ v3)}
            _ = v34
        }
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 int8 = 42
            v1 = (v1 * (^(v1)))
            if false {
                v37 := GenS4{F0: v4, F1: 42}
                _ = v37
            } else {
                v38 := v5["foo bar"]
                var v39 uint = 65535
                var v40 int8 = (v36 + (-(v36)))
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi35
            _ = v36
        }
        _ = fi31
    }
    fmt.Println(v6)
    v2 += (^(v2))
    for fi41 := 0; fi41 < 2; fi41++ {
        var v42 int = fi41
        fmt.Println(v2)
        v43 := GenS3{F0: (32767 > 42), F1: v3}
        _ = fi41
        _ = v42
        _ = v43
    }
    switch ("foo bar" + "hello") {
    case "":
        delete(v5, "foo bar")
    case "hello":
        for fi44 := 0; fi44 < 15; fi44++ {
            v45 := GenS5{F0: v4, F1: 2}
            v3 -= v3
            var v46 string = "foo bar"
            _ = fi44
            _ = v45
            _ = v46
        }
    case "\n\t":
        if ((!true) || false) {
            var v47 string = ""
            _ = v47
        } else {
            v2 = (v2 >> ((v4) & 63))
        }
    }
    v48 := GenS2{F0: -775, F1: 0}
    var v49 int32 = 42
    var v50 int32 = (v49 * (v49 / ((int32(v2)) | 1)))
    if false {
        v50 = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
