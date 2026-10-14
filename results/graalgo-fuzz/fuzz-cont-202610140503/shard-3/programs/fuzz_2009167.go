package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []string{"foo bar", "hello", "\n\t", "\n\t"}
    var v2 int64 = -1
    var v3 int = -729
    fmt.Println(v2, v2)
    v3 -= 42
    v4 := GenS1{F0: "\n\t"}
    for fi5 := 0; fi5 < 19; fi5++ {
        v3 = 32768
        _ = fi5
    }
    var v6 string = "a"
    if true {
        if len(v1) > 0 {
            v1[0] = ("foo bar" + v6)
        }
    } else {
        if (-2.25 != 1.5) {
            for fi7 := 0; fi7 < 9; fi7++ {
                switch "foo bar" {
                case "":
                    v1 = append(v1, (v6 + "a"))
                case "a":
                    fmt.Println(v3, fi7)
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v8 uint16 = 255
                _ = fi7
                _ = v8
            }
            v2 += v2
        } else {
            for fi9 := 0; fi9 < 15; fi9++ {
                if (-1 >= (v3 / (((v3 | v3)) | 1))) {
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 int16 = 1
                        v2 = (v2 ^ (v2 | (^(v2))))
                        var v12 int32 = 300
                        if true {
                            fmt.Println(v6, v11)
                            var v13 int64 = 42
                            v14 := GenS1{F0: ((v6 + v6) + v6)}
                            _ = v13
                            _ = v14
                        } else {
                            if true {
                                if false {
                                    _ = len(v1)
                                    v3 += fi9
                                }
                                v2 -= 0
                                switch (("foo bar" + ("a" + ("hello" + "\n\t"))) + ((v6 + v6) + "hello")) {
                                case "":
                                    var v15 int = (^(fi9))
                                    _ = v15
                                case "foo bar":
                                    v1 = append(v1, v6)
                                    v6 = v6
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v12 = (v12 % ((2147483647) | 1))
                                var v16 float32 = 1e10
                                v2 = v2
                                var v17 int = (v3 & fi9)
                                _ = v16
                                _ = v17
                            } else {
                                v3 = (fi10 & 274)
                            }
                            fmt.Println(v6, v2, fi10)
                        }
                        v1 = append(v1, ("" + v6))
                        if true {
                            v6 = v6
                        }
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    for fi18 := 0; fi18 < 5; fi18++ {
                        switch v6 {
                        case "hello":
                            v3 = v3
                        case "":
                            v6 = ("hello" + v6)
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (false || true) {
                            if ("hello" >= v6) {
                                if len(v1) > 0 {
                                    v1[0] = v6
                                }
                                var v19 int64 = 9223372036854775807
                                switch "foo bar" {
                                case "hello":
                                    var v20 string = v6
                                    _ = v20
                                case "foo bar":
                                    var v21 float32 = -2.25
                                    _ = v21
                                }
                                if false {
                                    fmt.Println(v19, v3, v19)
                                    fmt.Println(v19)
                                    v22 := GenS1{F0: v6}
                                    _ = v22
                                }
                                v1 = append(v1, "a")
                                v19 = 307
                                v23 := GenS1{F0: v6}
                                _ = v19
                                _ = v23
                            } else {
                                v1 = append(v1, v6)
                                var v24 uint64 = 2
                                _ = v24
                            }
                        }
                        v2 += v2
                        _ = fi18
                    }
                } else {
                    fmt.Println(v6)
                }
                v2 += (v2 << ((42) & 63))
                fmt.Println(fi9, v6)
                _ = fi9
            }
            v6 = (("hello" + (v6 + ((("\n\t" + "a") + v6) + v6))) + (("" + ("a" + (("hello" + (v6 + v6)) + v6))) + v6))
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v25 := GenS2{F0: 449}
    switch (^(v3)) {
    case -859:
        var v26 uint16 = 255
        _ = v26
    case 4294967295:
        v27 := GenS1{F0: "hello"}
        _ = v27
    }
    switch (v6 + v6) {
    case "foo bar":
        if true {
            v1 = append(v1, v6)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "":
        var v28 int64 = (v2 - int64(v3))
        _ = v28
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        if false {
            var v29 uint8 = 95
            var v30 int64 = v2
            _ = v29
            _ = v30
        } else {
            v6 = v6
            v3 += v3
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
