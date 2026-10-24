package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 127
    v3 := []string{"foo bar"}
    var v4 int = 9223372036854775807
    v5 := []int{v4, (v4 ^ (^(v4))), -2147483648}
    fmt.Println(v4, v4)
    if ("foo bar" == "hello") {
        if (-888 == -32768) {
            v6 := GenS1{F0: 2}
            _ = v6
        }
        var v7 uint64 = 953
        v2 = -9223372036854775808
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if ("\n\t" <= "") {
            var v9 string = ("a" + "foo bar")
            v1[fi8] = v9
            v3 = append(v3, v9)
            _ = v9
        } else {
            if (!(-73 == 127)) {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v4 = fi8
                    _ = fi10
                }
            }
            fmt.Println(fi8, fi8)
            for fi11 := 0; fi11 < 18; fi11++ {
                v2 = v2
                _ = fi11
            }
        }
        delete(v1, v4)
        _ = fi8
    }
    v2 = 769
    v12 := GenS1{F0: 286}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int32 = 2147483647
        v15 := GenS1{F0: 18446744073709551615}
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 = v2
    switch v4 {
    case 9223372036854775807:
        if false {
            v4 = (v4 % ((-683) | 1))
            v2 = (v2 + v2)
        }
    default:
        v4 += v4
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        switch 783 {
        case 89:
            v2 = 1
        case 32768:
            for fi17 := 0; fi17 < 15; fi17++ {
                v5 = append(v5, -2147483648)
                v4 += fi16
                var v18 int8 = 42
                _ = fi17
                _ = v18
            }
        }
        if (42 != -1) {
            for fi19 := 0; fi19 < 13; fi19++ {
                v2 = v2
                _ = fi19
            }
        }
        var v20 uint8 = 127
        _ = fi16
        _ = v20
    }
    switch (v4 >> ((460) & 63)) {
    case 9223372036854775807:
        if (true || false) {
            var v21 float32 = -2.25
            _ = v21
        } else {
            v2 += -1
            v22 := GenS1{F0: 42}
            var v23 int16 = 0
            _ = v22
            _ = v23
        }
    default:
        switch ("\n\t" + ("hello" + ("" + "a"))) {
        case "hello":
            var v24 string = "foo bar"
            _ = v24
        default:
            switch "foo bar" {
            case "foo bar":
                var v25 int64 = v2
                _ = v25
            case "a":
                for fi26 := 0; fi26 < 12; fi26++ {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 int8 = 0
                        if (true && false) {
                            switch (fi27 * fi27) {
                            case -9223372036854775808:
                                _ = len(v5)
                            default:
                                v29 := GenS1{F0: 1}
                                _ = v29
                            }
                        } else {
                            _ = len(v1)
                        }
                        var v30 string = "hello"
                        _ = fi27
                        _ = v28
                        _ = v30
                    }
                    _ = fi26
                }
            }
        }
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 int32 = 42
        var v33 bool = (!(3.14159 <= 3.14159))
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        v35 := GenS1{F0: 2}
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
