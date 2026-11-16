package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 uint8 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint = 2
    var v2 int8 = 0
    var v3 int = 9223372036854775807
    var v4 int64 = -1
    var v5 int8 = v2
    v6 := GenS1{F0: "\n\t", F1: 808}
    if ((float32(1e10) * (-2.25 * 3.14159)) <= float32(3.14159)) {
        v7 := GenS2{F0: 595}
        v8 := GenS2{F0: 0}
        var v9 uint32 = 1
        _ = v7
        _ = v8
        _ = v9
    } else {
        v4 = (v4 & -95)
        var v10 int = v3
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v3, v3, v5)
        case "hello":
            var v11 uint32 = 1
            _ = v11
        case "\n\t":
            for fi12 := 0; fi12 < 7; fi12++ {
                switch v10 {
                case 42:
                    var v13 uint = (v1 + v1)
                    _ = v13
                case 4294967295:
                    v14 := GenS3{F0: true, F1: (fi12 >> (((v1 % ((42) | 1))) & 63)), F2: 1}
                    _ = v14
                default:
                    if false {
                        var v15 uint32 = 42
                        _ = v15
                    } else {
                        v2 -= (v2 % (((v5 + 127)) | 1))
                        v16 := []int{v3, (fi12 << ((170) & 63)), (fi12 | (fi12 - (v3 + (-(v3)))))}
                        for fi17 := 0; fi17 < 1; fi17++ {
                            var v18 int8 = (v2 & v5)
                            if len(v16) > 0 {
                                v16[0] = -2147483648
                            }
                            var v19 float32 = 1.5
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        _ = v16
                    }
                }
                v20 := []int{(v3 / ((v3) | 1)), v10, fi12, (v3 - v3)}
                _ = fi12
                _ = v20
            }
        default:
            var v21 bool = ("" < "foo bar")
            _ = v21
        }
        _ = v10
    }
    fmt.Println(v3)
    for fi22 := 0; fi22 < 11; fi22++ {
        v1 = (v1 >> ((v1) & 63))
        _ = fi22
    }
    switch (v3 << (((v1 << ((1) & 63))) & 63)) {
    case 748:
        v2 -= v2
    case 814:
        v23 := GenS1{F0: ("\n\t" + "\n\t"), F1: 32767}
        _ = v23
    case -312:
        if (float32(((((1.5 - (float64(float32((1.5 * 3.14159))) - 1e10)) - 1.5) + float64((float32(1e10) + 0.0))) - ((3.14159 + (1e10 - 3.14159)) + 3.14159))) != 1e10) {
            v24 := []int{70}
            if true {
                v5 = v2
            }
            if (false && (!(!false))) {
                var v25 string = "\n\t"
                for fi26 := 0; fi26 < 16; fi26++ {
                    for fi27 := 0; fi27 < 13; fi27++ {
                        v25 = (v25 + v25)
                        _ = len(v24)
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        _ = fi27
                    }
                    _ = fi26
                }
                _ = v25
            }
            _ = v24
        } else {
            v28 := GenS2{F0: 4294967295}
            _ = v28
        }
    default:
        if false {
            v29 := GenS1{F0: "foo bar", F1: 863}
            switch ("\n\t" + ("" + "\n\t")) {
            case "foo bar":
                v30 := GenS2{F0: 0}
                _ = v30
            }
            _ = v29
        } else {
            for fi31 := 0; fi31 < 14; fi31++ {
                if ((v4 % ((-488) | 1)) == -1) {
                    for fi32 := 0; fi32 < 1; fi32++ {
                        switch "" {
                        case "\n\t":
                            var v33 bool = (0.1 == ((1e10 * float32(0.0)) + (((0.1 * 1e10) + 1.5) * float32(3.14159))))
                            _ = v33
                        }
                        var v34 uint64 = 0
                        _ = fi32
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 12; fi35++ {
                        v3 = v3
                        _ = fi35
                    }
                } else {
                    fmt.Println(v2, v1)
                }
                _ = fi31
            }
            if (!false) {
                for fi36 := 0; fi36 < 5; fi36++ {
                    var v37 uint16 = 2
                    _ = fi36
                    _ = v37
                }
            } else {
                var v38 float64 = -2.25
                v39 := []int{(v3 * (v3 | v3)), v3}
                switch (v3 - v3) {
                case -1:
                    var v40 uint32 = 2
                    _ = v40
                case 32767:
                    if false {
                        var v41 uint32 = 2147483647
                        var v42 uint8 = 42
                        if (v4 > v4) {
                            v5 -= (v2 & -1)
                            v41 = 42
                            v2 -= v5
                        }
                        _ = v41
                        _ = v42
                    }
                case -9223372036854775808:
                    var v43 bool = (v38 < 1.5)
                    _ = v43
                }
                _ = v38
                _ = v39
            }
            var v44 float32 = -2.25
            _ = v44
        }
    }
    v45 := GenS2{F0: 206}
    v46 := []int{0, -9223372036854775808, v3}
    v3 = -2147483648
    if len(v46) > 0 {
        v46[0] = v3
    }
    fmt.Println(v5, v5)
    if (!true) {
        if (84 > 128) {
            if len(v46) > 0 {
                v46[0] = 1
            }
            if (!(!true)) {
                switch ((("foo bar" + "\n\t") + "\n\t") + "") {
                case "hello":
                    var v47 int16 = 898
                    _ = v47
                case "a":
                    fmt.Println(v2, v1, v4)
                }
            } else {
                _ = len(v46)
                v48 := GenS4{F0: 1}
                _ = v48
            }
            v46 = append(v46, -1)
        }
    } else {
        v3 -= (v3 << ((65535) & 63))
        v49 := GenS2{F0: 14}
        _ = v49
    }
    for fi50 := 0; fi50 < 7; fi50++ {
        var v51 uint = 1
        _ = fi50
        _ = v51
    }
    fmt.Println(v1, v5, v1)
    for fi52 := 0; fi52 < 2; fi52++ {
        var v53 int32 = 1
        v54 := GenS1{F0: "foo bar", F1: 32767}
        _ = fi52
        _ = v53
        _ = v54
    }
    if len(v46) >= 2 {
        v46 = v46[1:len(v46)]
    }
    v55 := GenS1{F0: ("\n\t" + ""), F1: 792}
    v2 = (v5 / ((v2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v55
    _ = v6
}
