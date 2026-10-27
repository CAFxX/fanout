package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 int32 = 0
    v2 := []string{"foo bar", "hello", ("hello" + ("hello" + "hello")), "hello"}
    v3 := make(map[int]string)
    var v4 float64 = (-2.25 * 3.14159)
    v5 := make(map[int]string)
    var v6 int = 32767
    v4 += v4
    fmt.Println(v4, v6, v4)
    switch "\n\t" {
    case "a":
        if (float64(3.14159) < v4) {
            for fi7 := 0; fi7 < 3; fi7++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi7
            }
            if false {
                v6 = 2147483647
                v2 = append(v2, "\n\t")
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                if false {
                    fmt.Println(v6, v6)
                } else {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    var v9 uint = 42
                    var v10 uint16 = 255
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
        }
    case "\n\t":
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float32 = 1.5
            if (42 >= 9223372036854775807) {
                var v13 int64 = 0
                _ = v13
            }
            var v14 uint16 = 746
            _ = fi11
            _ = v12
            _ = v14
        }
    case "hello":
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: 42}
            _ = fi15
            _ = v16
        }
    }
    var v17 uint32 = 42
    _ = len(v3)
    v18 := GenS1{F0: 42}
    fmt.Println(v4)
    if true {
        if (0 == -1) {
            var v19 int = v6
            var v20 uint64 = 1
            var v21 uint64 = (v20 / ((v20) | 1))
            _ = v19
            _ = v20
            _ = v21
        }
        _ = len(v5)
        if (!(false || false)) {
            if (!false) {
                var v22 float64 = 1.5
                _ = v22
            } else {
                for fi23 := 0; fi23 < 20; fi23++ {
                    for fi24 := 0; fi24 < 15; fi24++ {
                        var v25 uint32 = (^(v17))
                        if (!(true && ((1 >= 4294967295) && (-394 <= -1)))) {
                            _ = len(v2)
                        } else {
                            fmt.Println(v17, fi24, v6)
                            fmt.Println(v6, v25)
                            fmt.Println(v4, fi23)
                            switch 32767 {
                            case 127:
                                v17 = 1
                            case 42:
                                if (42 <= 127) {
                                    var v26 uint16 = 42
                                    _ = v26
                                } else {
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                    v1 = v1
                                }
                            }
                        }
                        v17 -= v17
                        _ = fi24
                        _ = v25
                    }
                    _ = fi23
                }
            }
        }
    }
    switch v6 {
    case -320:
        var v27 string = ("a" + ("a" + "a"))
        _ = v27
    case -9223372036854775808:
        v6 -= v6
    default:
        v28 := GenS1{F0: 42}
        _ = v28
    }
    v4 = v4
    switch v6 {
    case -20:
        var v29 string = ("\n\t" + "hello")
        _ = v29
    case 42:
        switch v6 {
        case 32768:
            v30 := GenS1{F0: 128}
            _ = v30
        }
    }
    var v31 float64 = v4
    for fi32 := 0; fi32 < 13; fi32++ {
        var v33 int = v6
        var v34 uint16 = 32767
        _ = fi32
        _ = v33
        _ = v34
    }
    v6 = 42
    v4 += 0.1
    var v35 bool = (int64(v1) <= -1)
    v2 = append(v2, ("" + "foo bar"))
    for fi36 := 0; fi36 < 4; fi36++ {
        v35 = (!((!v35) && v35))
        v3[v6] = "a"
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
