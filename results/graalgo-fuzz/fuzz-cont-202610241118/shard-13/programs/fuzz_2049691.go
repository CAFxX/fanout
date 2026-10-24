package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []string{("\n\t" + "hello"), ("a" + "foo bar")}
    v2 := make(map[string]bool)
    var v3 bool = true
    _ = len(v2)
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 0, F1: 32767}
        _ = fi4
        _ = v5
    }
    fmt.Println(v3)
    v6 := GenS1{F0: 2147483647, F1: 4294967295}
    fmt.Println(v3)
    switch ("hello" + "\n\t") {
    case "\n\t":
        if (((!(!true)) && (623 > 42)) || v3) {
            v3 = (!((1000 < 1000) && (2 == 2)))
            if (!false) {
                for fi7 := 0; fi7 < 2; fi7++ {
                    v3 = (!(42 == -193))
                    if (1 == 104) {
                        if true {
                            v3 = (2 < 157)
                            var v8 string = ""
                            v9 := GenS1{F0: 42, F1: fi7}
                            _ = v8
                            _ = v9
                        } else {
                            var v10 int32 = 668
                            v3 = (!(((-128 != -128) || (v3 || ((1 > 387) || v3))) || (-2.25 != 1.5)))
                            _ = v10
                        }
                        var v11 uint = 4294967295
                        if (2147483647 < -968) {
                            if v3 {
                                v3 = v3
                            }
                            if len(v1) > 0 {
                                v1[0] = ("hello" + ("a" + "\n\t"))
                            }
                            v3 = (-1 > 113)
                        } else {
                            v11 -= v11
                            if len(v1) > 0 {
                                v1[0] = (("\n\t" + "hello") + "foo bar")
                            }
                            v3 = v3
                        }
                        _ = v11
                    }
                    _ = fi7
                }
            } else {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := GenS1{F0: 2147483648, F1: fi12}
                    _ = len(v1)
                    v14 := GenS1{F0: 2147483647, F1: fi12}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 3; fi15++ {
                    if false {
                        var v16 int = (fi15 % ((-879) | 1))
                        if false {
                            v3 = ("" > "\n\t")
                            v16 = fi15
                            var v17 int8 = 1
                            _ = v17
                        } else {
                            v18 := GenS2{F0: 331, F1: 821, F2: 199}
                            var v19 string = "a"
                            v16 = -461
                            _ = v18
                            _ = v19
                        }
                        var v20 uint8 = 127
                        _ = v16
                        _ = v20
                    }
                    if (((-(fi15)) >= fi15) || ((-32768 >= -1) && v3)) {
                        switch "foo bar" {
                        case "a":
                            _ = len(v1)
                        case "foo bar":
                            v21 := GenS3{F0: -110}
                            _ = v21
                        case "":
                            var v22 uint = 255
                            _ = v22
                        }
                    } else {
                        var v23 int16 = -894
                        _ = v23
                    }
                    v3 = (649 <= 0)
                    _ = fi15
                }
            }
            if false {
                var v24 int64 = 42
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v1 = append(v1, "\n\t")
                _ = v24
            } else {
                v25 := v2[""]
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v25
            }
        } else {
            for fi26 := 0; fi26 < 9; fi26++ {
                v3 = v3
                if (!v3) {
                    var v27 int = fi26
                    _ = v27
                } else {
                    var v28 int8 = -1
                    _ = v28
                }
                _ = fi26
            }
        }
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        for fi30 := 0; fi30 < 11; fi30++ {
            v3 = v3
            if (v3 || v3) {
                v31 := GenS2{F0: 2, F1: 2147483648, F2: 2}
                _ = v31
            }
            v3 = (!v3)
            var v32 float32 = float32(1.5)
            var v33 uint8 = 0
            _ = fi30
            _ = v32
            _ = v33
        }
        _ = fi29
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        _ = len(v2)
        var v35 int64 = -567
        var v36 int64 = v35
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := GenS2{F0: 2, F1: 2147483648, F2: 356}
    v3 = v3
    v38 := GenS1{F0: 2147483648, F1: 127}
    switch 490 {
    case 9223372036854775807:
        v3 = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v6
}
