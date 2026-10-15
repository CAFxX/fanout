package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 string; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int32; F2 int16 }
func genFn0() string {
    for fi4 := 0; fi4 < 6; fi4++ {
        if (!(fi4 != 4294967295)) {
            for fi5 := 0; fi5 < 12; fi5++ {
                if false {
                    var v6 int = fi4
                    fmt.Println(v6, fi4)
                    switch (v6 >> ((42) & 63)) {
                    case 42:
                        if (0.0 < (float32((1e10 + 3.14159)) + 1.5)) {
                            v7 := GenS1{F0: 1, F1: 0}
                            var v8 int16 = -1
                            _ = v7
                            _ = v8
                        }
                    case 32768:
                        var v9 uint = 18446744073709551615
                        _ = v9
                    }
                    _ = v6
                } else {
                    var v10 uint = 489
                    if (false || false) {
                        v11 := []int{}
                        _ = v11
                    }
                    v12 := GenS2{F0: 42, F1: (v10 | (v10 | 42))}
                    _ = v10
                    _ = v12
                }
                _ = fi5
            }
            for fi13 := 0; fi13 < 14; fi13++ {
                fmt.Println(fi4, fi4)
                _ = fi13
            }
        }
        var v14 string = "\n\t"
        _ = fi4
        _ = v14
    }
    var v15 int = -334
    switch "hello" {
    case "foo bar":
        v15 += (v15 / (((-(v15))) | 1))
    case "hello":
        var v16 int32 = 42
        _ = v16
    case "a":
        v15 = (v15 * (^(v15)))
    default:
        var v17 int = (-(v15))
        _ = v17
    }
    _ = v15
    return ""
}

func main() {
    v1 := []int{32767, 0, 128}
    v2 := make(map[string]string)
    var v3 int16 = 1000
    _ = len(v2)
    v18 := GenS3{F0: "a", F1: 42, F2: 980}
    if ("a" >= "\n\t") {
        switch -110 {
        case -989:
            if false {
                var v19 int8 = 42
                for fi20 := 0; fi20 < 15; fi20++ {
                    v21 := GenS3{F0: "hello", F1: 100, F2: 2}
                    v19 -= (v19 + (v19 * (v19 / (((v19 * (v19 & -70))) | 1))))
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        case 1:
            if ((4294967295 >= 32768) && true) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 uint = 4294967295
                    v23 -= v23
                    _ = fi22
                    _ = v23
                }
                fmt.Println(v3)
                fmt.Println(v3)
            }
        case -2147483648:
            _ = len(v2)
        }
    } else {
        var v24 int8 = 1
        fmt.Println(v24, v24)
        fmt.Println(v3, v3)
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        fmt.Println(fi25, v3)
        var v26 int16 = v3
        for fi27 := 0; fi27 < 20; fi27++ {
            v28 := GenS3{F0: "a", F1: 127, F2: 4294967295}
            switch "\n\t" {
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "foo bar":
                v3 = 1000
                var v29 int8 = 0
                _ = v29
            case "":
                var v30 uint16 = 921
                _ = v30
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch ("hello" + ("" + "\n\t")) {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = int(v26)
                }
            default:
                fmt.Println(v26)
            }
            _ = fi27
            _ = v28
        }
        _ = fi25
        _ = v26
    }
    v3 = (v3 << ((255) & 63))
    if true {
        var v31 uint = 42
        for fi32 := 0; fi32 < 20; fi32++ {
            v31 = 18446744073709551615
            v33 := GenS3{F0: "", F1: 0, F2: 535}
            var v34 int32 = 1
            _ = fi32
            _ = v33
            _ = v34
        }
        v35 := GenS4{F0: uint8(v31)}
        _ = v31
        _ = v35
    }
    if (42 == 1) {
        v3 = v3
        v36 := GenS5{F0: (v3 / (((v3 ^ 1000)) | 1)), F1: -1, F2: v3}
        v3 += v3
        _ = v36
    }
    delete(v2, (("hello" + "hello") + ("hello" + ((("hello" + "foo bar") + "\n\t") + "foo bar"))))
    if len(v1) > 0 {
        v1[0] = 127
    }
    _ = len(v2)
    delete(v2, "\n\t")
    if true {
        var v37 int64 = 0
        v3 -= v3
        var v38 float32 = 3.14159
        _ = v37
        _ = v38
    }
    var v39 int64 = -682
    for fi40 := 0; fi40 < 12; fi40++ {
        var v41 int16 = -918
        _ = fi40
        _ = v41
    }
    switch (("foo bar" + ("hello" + "a")) + "hello") {
    case "a":
        var v42 uint = 65535
        _ = v42
    case "foo bar":
        var v43 uint16 = 32767
        _ = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v39
}
