package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int64 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 string; F2 uint8 }

func main() {
    var v1 float64 = float64(3.14159)
    var v2 int64 = -1
    v3 := make(map[string]bool)
    switch "\n\t" {
    case "":
        v4 := []int{-2147483648, 431, 2147483647}
        _ = v4
    case "foo bar":
        if (1 < 0) {
            fmt.Println(v2)
            for fi5 := 0; fi5 < 10; fi5++ {
                v2 += 1
                _ = fi5
            }
        } else {
            v6 := GenS1{F0: 0, F1: "\n\t", F2: (v2 << ((42) & 63))}
            _ = v6
        }
    case "\n\t":
        var v7 int8 = 127
        _ = v7
    }
    fmt.Println(v2, v2)
    if (677 < 753) {
        if (!(101 == 65535)) {
            var v8 int32 = 0
            v1 -= 1e10
            v8 -= v8
            _ = v8
        } else {
            if true {
                var v9 uint64 = 162
                _ = v9
            } else {
                if false {
                    v2 = v2
                } else {
                    if false {
                        v3[("a" + "hello")] = (772 >= 717)
                    } else {
                        v10 := []int{-9223372036854775808, 32768, 451}
                        var v11 float32 = 1e10
                        _ = v10
                        _ = v11
                    }
                }
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        v12 := GenS1{F0: v2, F1: ("\n\t" + (("a" + "\n\t") + "hello")), F2: -1}
                        _ = v12
                    case "a":
                        var v13 uint64 = 1
                        _ = v13
                    }
                    var v14 uint = 1
                    _ = v14
                }
            }
        }
    }
    var v15 int64 = v2
    if false {
        v15 += (v2 & (^(v2)))
        var v16 int64 = -563
        for fi17 := 0; fi17 < 3; fi17++ {
            switch ("a" + "a") {
            case "a":
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 string = "hello"
                    v15 -= v16
                    _ = fi18
                    _ = v19
                }
            case "\n\t":
                switch (fi17 & (fi17 >> ((286) & 63))) {
                case -2147483648:
                    fmt.Println(v15, v1, fi17)
                }
            case "foo bar":
                if (684 <= 0) {
                    v20 := []int{fi17, (fi17 * -306), (fi17 / (((fi17 + fi17)) | 1)), (fi17 >> ((4294967295) & 63))}
                    if true {
                        var v21 uint32 = 4294967295
                        if ((2147483647 < 861) || true) {
                            var v22 bool = true
                            _ = v22
                        } else {
                            var v23 float64 = v1
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = v23
                        }
                        _ = v21
                    }
                    for fi24 := 0; fi24 < 8; fi24++ {
                        var v25 uint8 = 127
                        _ = fi24
                        _ = v25
                    }
                    _ = v20
                } else {
                    var v26 uint = 1
                    var v27 uint16 = 65535
                    _ = v26
                    _ = v27
                }
            }
            var v28 uint16 = 42
            if false {
                var v29 uint = 1
                _ = v29
            }
            _ = fi17
            _ = v28
        }
        _ = v16
    }
    v2 -= (v2 | v15)
    v15 -= 1
    if false {
        var v30 float32 = (-2.25 * 0.1)
        var v31 bool = ("foo bar" >= (("a" + ("\n\t" + ("\n\t" + ("foo bar" + "a")))) + ("hello" + "hello")))
        var v32 uint = 255
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 float32 = -2.25
    v34 := GenS1{F0: (v2 ^ 0), F1: ("foo bar" + "\n\t"), F2: (v15 ^ v2)}
    var v35 int64 = v15
    v36 := GenS2{F0: ("\n\t" + "\n\t"), F1: (160 < 255)}
    switch 4294967295 {
    case -274:
        if (((1.5 + v1) + 0.0) < v1) {
            v37 := GenS1{F0: v35, F1: "a", F2: (v2 & 0)}
            var v38 uint64 = 4294967295
            if false {
                var v39 int = 127
                var v40 uint16 = 255
                _ = v39
                _ = v40
            }
            _ = v37
            _ = v38
        }
    case -486:
        for fi41 := 0; fi41 < 9; fi41++ {
            v42 := GenS3{F0: (fi41 / ((fi41) | 1)), F1: ((("" + "foo bar") + "") + ("foo bar" + (("a" + "") + "a"))), F2: 2}
            switch "foo bar" {
            case "":
                var v43 int32 = 418
                _ = v43
            case "hello":
                var v44 string = "hello"
                _ = v44
            case "a":
                for fi45 := 0; fi45 < 11; fi45++ {
                    var v46 uint = 4294967295
                    v3["\n\t"] = true
                    var v47 uint64 = 897
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            _ = fi41
            _ = v42
        }
    }
    v1 = -2.25
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
