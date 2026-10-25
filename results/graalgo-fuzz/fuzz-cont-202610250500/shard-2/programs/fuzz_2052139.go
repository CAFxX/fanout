package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 int32; F2 int32 }

func main() {
    var v1 float32 = -2.25
    v2 := []string{"foo bar", "a", "hello", "hello"}
    v3 := make(map[int]bool)
    var v4 int8 = 122
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            switch fi6 {
            case -1:
                fmt.Println(v1, v4)
            case -256:
                var v7 int16 = -32768
                _ = v7
            case 812:
                v8 := GenS1{F0: 92, F1: 0}
                _ = v8
            }
            var v9 uint = 1
            _ = fi6
            _ = v9
        }
        _ = fi5
    }
    v4 = v4
    v1 = 0.1
    v10 := GenS1{F0: 42, F1: 686}
    v4 += v4
    v11 := GenS2{F0: 2}
    var v12 int32 = -2147483648
    var v13 uint = 255
    switch ("\n\t" + "hello") {
    case "a":
        if false {
            if true {
                fmt.Println(v12, v1)
                if (!(!true)) {
                    switch 2147483647 {
                    case 188:
                        switch ("hello" + (("foo bar" + "foo bar") + (((("a" + ((("" + ("" + "")) + ("foo bar" + "foo bar")) + "\n\t")) + (("a" + "hello") + ("foo bar" + "foo bar"))) + "") + "a"))) {
                        case "":
                            var v14 uint64 = 0
                            _ = v14
                        case "foo bar":
                            fmt.Println(v13)
                        default:
                            for fi15 := 0; fi15 < 10; fi15++ {
                                for fi16 := 0; fi16 < 13; fi16++ {
                                    v4 = 127
                                    v3[(fi15 % (((fi16 ^ fi15)) | 1))] = (1e10 >= v1)
                                    if (873 < -1) {
                                        var v17 int16 = -802
                                        fmt.Println(v4, v1, fi16)
                                        var v18 uint16 = 976
                                        switch (^(fi16)) {
                                        case -577:
                                            if ((!(0 < 41)) && true) {
                                                v3[(^(fi16))] = (!false)
                                                fmt.Println(v1)
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                                if false {
                                                    var v19 uint8 = 127
                                                    v13 = (v13 ^ v13)
                                                    var v20 int16 = v17
                                                    _ = v19
                                                    _ = v20
                                                } else {
                                                    v21 := GenS3{F0: 2, F1: -878}
                                                    v13 = v13
                                                    var v22 int32 = (^(v12))
                                                    _ = v21
                                                    _ = v22
                                                }
                                            }
                                        case 1:
                                            v12 = 2147483647
                                        case -2147483648:
                                            _ = len(v2)
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v23 float32 = v1
                                        var v24 uint64 = 529
                                        if len(v2) > 0 {
                                            v2[0] = "hello"
                                        }
                                        _ = v23
                                        _ = v24
                                    }
                                    fmt.Println(fi15, v1)
                                    _ = fi16
                                }
                                fmt.Println(v4, v4, v12)
                                v25 := GenS4{F0: 114}
                                _ = fi15
                                _ = v25
                            }
                        }
                    case 2147483647:
                        var v26 int = 127
                        _ = v26
                    }
                    switch 2147483647 {
                    case 609:
                        for fi27 := 0; fi27 < 11; fi27++ {
                            switch (fi27 >> ((v13) & 63)) {
                            case 42:
                                var v28 int = -2147483648
                                _ = v28
                            }
                            _ = fi27
                        }
                    case -2147483648:
                        v29 := GenS2{F0: 484}
                        _ = v29
                    default:
                        var v30 uint64 = 98
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 14; fi31++ {
                        var v32 uint8 = 1
                        _ = fi31
                        _ = v32
                    }
                }
            } else {
                var v33 int64 = 9223372036854775807
                v1 = v1
                _ = v33
            }
        }
    }
    if true {
        for fi34 := 0; fi34 < 9; fi34++ {
            v4 += v4
            _ = fi34
        }
        if false {
            var v35 uint64 = 1
            _ = v35
        } else {
            for fi36 := 0; fi36 < 1; fi36++ {
                if (true && ((!(699 == -1)) && true)) {
                    if ((v12 << ((v13) & 63)) <= (v12 / ((v12) | 1))) {
                        var v37 uint = (v13 ^ 1)
                        for fi38 := 0; fi38 < 11; fi38++ {
                            v39 := GenS5{F0: -452, F1: v12, F2: (^(v12))}
                            _ = fi38
                            _ = v39
                        }
                        var v40 uint64 = 0
                        _ = v37
                        _ = v40
                    }
                    var v41 string = ""
                    _ = v41
                } else {
                    if (fi36 > (fi36 & fi36)) {
                        if (false && true) {
                            fmt.Println(fi36, v13, fi36)
                            v1 = v1
                            var v42 int16 = 100
                            _ = v42
                        }
                    } else {
                        var v43 float32 = 3.14159
                        _ = v43
                    }
                }
                v44 := GenS5{F0: 939, F1: (v12 / (((v12 << ((v13) & 63))) | 1)), F2: v12}
                _ = fi36
                _ = v44
            }
        }
    }
    fmt.Println(v1, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
