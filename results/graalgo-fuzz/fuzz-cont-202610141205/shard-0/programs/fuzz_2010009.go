package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int }
func genFn0(p0 bool, p1 int64) uint32 {
    var v7 uint = 65535
    v8 := GenS1{F0: "\n\t"}
    p0 = (32767 >= 976)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 2, 857}
    var v2 uint64 = 1
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    var v5 int8 = 9
    var v6 int8 = -82
    v9 := GenS2{F0: v5}
    if (2147483647 <= 281) {
        var v10 uint16 = 42
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := GenS2{F0: (v5 + (^(v6)))}
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v10 = (v10 * (v10 | v10))
        _ = v10
    } else {
        var v14 int = 9223372036854775807
        fmt.Println(v5)
        _ = v14
    }
    v15 := GenS3{F0: 2147483648}
    v16 := GenS1{F0: ("foo bar" + "hello")}
    fmt.Println(v6, v5, v2)
    if false {
        v1 = append(v1, 240)
        if true {
            var v17 uint32 = 0
            for fi18 := 0; fi18 < 19; fi18++ {
                _ = len(v4)
                v19 := GenS1{F0: "a"}
                v2 = (^(v2))
                _ = fi18
                _ = v19
            }
            v20 := GenS4{F0: 148, F1: v6}
            _ = v17
            _ = v20
        }
        var v21 string = ""
        _ = v21
    } else {
        v22 := GenS5{F0: "\n\t"}
        _ = v22
    }
    v5 -= (v6 + (v6 >> ((65535) & 63)))
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint = 1
        switch fi23 {
        case 32767:
            for fi25 := 0; fi25 < 17; fi25++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch (-(fi25)) {
                case 1:
                    _ = len(v1)
                    _ = len(v1)
                    v24 += (^(v24))
                case 0:
                    v4[851] = 1
                case -1:
                    _ = len(v4)
                default:
                    v2 = v2
                }
                if (0.1 < -2.25) {
                    v26 := v3[("a" + "\n\t")]
                    _ = v26
                } else {
                    fmt.Println(v6, fi23, v2)
                    switch ("\n\t" + ("foo bar" + "a")) {
                    case "\n\t":
                        var v27 int32 = -2147483648
                        _ = v27
                    case "hello":
                        switch "a" {
                        case "\n\t":
                            v2 = 42
                        case "hello":
                            var v28 int16 = 42
                            v5 = (v6 + (v6 - (v5 | (v5 % (((v6 ^ v6)) | 1)))))
                            if false {
                                v28 = (v28 + v28)
                            } else {
                                v29 := GenS6{F0: fi25}
                                delete(v4, 2147483647)
                                var v30 int32 = 0
                                _ = v29
                                _ = v30
                            }
                            _ = v28
                        default:
                            var v31 string = ""
                            _ = v31
                        }
                    default:
                        switch "hello" {
                        case "hello":
                            v6 = (v5 | (v6 - (v5 - 1)))
                        }
                    }
                    v2 = v2
                }
                if (32767 < 32768) {
                    delete(v3, "a")
                    var v32 string = ""
                    var v33 uint = (v24 - (^(v24)))
                    v33 = (v24 * (v24 / (((^(v24))) | 1)))
                    _ = v32
                    _ = v33
                }
                v6 = (^(v5))
                v34 := GenS5{F0: (("foo bar" + "hello") + ((("hello" + "") + "a") + "a"))}
                var v35 uint = (v24 + 1)
                _ = fi25
                _ = v34
                _ = v35
            }
        case -1:
            v24 = 255
        default:
            if len(v1) > 0 {
                v1[0] = 65535
            }
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v6, v2, v2)
    if true {
        for fi36 := 0; fi36 < 20; fi36++ {
            v37 := GenS4{F0: 0, F1: v6}
            for fi38 := 0; fi38 < 10; fi38++ {
                v39 := v3["a"]
                _ = fi38
                _ = v39
            }
            _ = fi36
            _ = v37
        }
        v40 := GenS3{F0: 2147483648}
        switch 128 {
        case 664:
            var v41 uint64 = v2
            _ = v41
        }
        _ = v40
    }
    v42 := v3["hello"]
    var v43 uint8 = 82
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v9
}
