package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 string; F1 int; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 string; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    var v4 bool = (false && ((0.1 * 1.5) != 3.14159))
    v5 := []int{127, 144, -9223372036854775808}
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            v3[2147483647] = ""
            if v4 {
                v1 += (v1 - 1)
                var v8 int8 = 0
                if true {
                    var v9 uint64 = 698
                    if v4 {
                        switch ((("\n\t" + "\n\t") + "foo bar") + "foo bar") {
                        case "hello":
                            v10 := GenS1{F0: 42}
                            _ = v10
                        case "":
                            v9 = v9
                        case "a":
                            v1 += v1
                        default:
                            _ = len(v2)
                        }
                        var v11 string = "hello"
                        _ = v11
                    }
                    _ = v9
                }
                _ = v8
            } else {
                v12 := GenS1{F0: 100}
                v1 = 1
                _ = v12
            }
            if true {
                if (521 > 32767) {
                    if (!v4) {
                        if (!(!(42 == 65535))) {
                            v1 -= v1
                            var v13 float32 = 1.5
                            if true {
                                v13 -= float32(-2.25)
                                var v14 uint16 = 507
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v14
                            }
                            _ = v13
                        } else {
                            switch fi6 {
                            case 0:
                                delete(v3, 840)
                            default:
                                v15 := GenS2{F0: "\n\t", F1: 526}
                                _ = v15
                            }
                            v16 := GenS3{F0: "", F1: (^(fi7)), F2: 288}
                            v1 -= v1
                            _ = v16
                        }
                        v4 = (42 >= 0)
                    } else {
                        v4 = v4
                    }
                }
                v4 = false
                v4 = false
                var v17 int32 = -1
                _ = v17
            }
            var v18 int16 = 202
            v19 := GenS4{F0: -1, F1: 42}
            _ = fi7
            _ = v18
            _ = v19
        }
        _ = fi6
    }
    v20 := GenS5{F0: 1, F1: 0}
    v21 := GenS6{F0: "", F1: 42, F2: 32767}
    var v22 int16 = 32767
    if (v4 || true) {
        _ = len(v3)
        if false {
            var v23 int64 = 922
            if v4 {
                v24 := GenS2{F0: ("a" + "foo bar"), F1: 4294967295}
                switch 1 {
                case -1:
                    var v25 string = "a"
                    _ = v25
                }
                _ = v24
            } else {
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 int = fi26
                    v22 = 42
                    _ = fi26
                    _ = v27
                }
            }
            _ = v23
        } else {
            fmt.Println(v4)
        }
        var v28 int16 = v22
        _ = v28
    } else {
        switch (("\n\t" + ("\n\t" + ("a" + "a"))) + ("hello" + ("a" + ("hello" + "hello")))) {
        case "hello":
            v29 := GenS6{F0: ("hello" + "\n\t"), F1: 0, F2: 0}
            _ = v29
        }
    }
    v22 += v22
    fmt.Println(v4, v22, v4)
    fmt.Println(v1)
    for fi30 := 0; fi30 < 8; fi30++ {
        v31 := GenS5{F0: 42, F1: 1}
        v32 := GenS4{F0: 42, F1: 0}
        var v33 uint8 = v1
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    v1 = v1
    v34 := GenS5{F0: 4294967295, F1: 2147483647}
    var v35 uint64 = 549
    switch "\n\t" {
    case "hello":
        for fi36 := 0; fi36 < 17; fi36++ {
            for fi37 := 0; fi37 < 13; fi37++ {
                v22 -= (v22 - v22)
                fmt.Println(v1, fi36)
                _ = fi37
            }
            if true {
                for fi38 := 0; fi38 < 10; fi38++ {
                    switch "hello" {
                    case "foo bar":
                        var v39 string = ""
                        _ = v39
                    }
                    _ = fi38
                }
            }
            _ = fi36
        }
    case "a":
        v40 := GenS3{F0: ("\n\t" + "\n\t"), F1: 0, F2: 386}
        _ = v40
    default:
        if v4 {
            var v41 int = 42
            v42 := GenS1{F0: 255}
            v35 += (v35 | v35)
            _ = v41
            _ = v42
        }
    }
    v43 := GenS7{F0: (v1 % ((0) | 1))}
    var v44 uint8 = v1
    var v45 uint64 = v35
    v1 += (v44 / ((v44) | 1))
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
