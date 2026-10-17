package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -633
    var v3 int8 = 127
    v4 := []string{"foo bar", "a", "hello"}
    if true {
        v3 -= int8(v2)
        fmt.Println(v2)
    } else {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
    }
    v2 = 1
    _ = len(v4)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 string = ""
        _ = fi5
        _ = v6
    }
    v4 = append(v4, "hello")
    v7 := v1[2]
    v7 = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            switch "" {
            case "":
                var v10 int16 = -857
                _ = v10
            case "hello":
                var v11 int64 = -1
                _ = v11
            }
            v3 = -50
            _ = fi9
        }
        var v12 uint = 249
        fmt.Println(v7, v3)
        _ = fi8
        _ = v12
    }
    var v13 int32 = -2147483648
    if (1e10 != float64((0.1 + 1.5))) {
        v7 = (v7 + (v7 << ((4294967295) & 63)))
    } else {
        v3 += (v3 % ((30) | 1))
        switch "hello" {
        case "a":
            var v14 uint32 = (v7 % (((^(v7))) | 1))
            _ = v14
        case "foo bar":
            v3 += int8(v13)
        case "hello":
            for fi15 := 0; fi15 < 10; fi15++ {
                switch (v2 - (v2 - -591)) {
                case -149:
                    switch (("\n\t" + "") + ("hello" + "foo bar")) {
                    case "hello":
                        if ((!true) && (!false)) {
                            if true {
                                var v16 float64 = -2.25
                                v17 := GenS1{F0: 772}
                                v4 = append(v4, "a")
                                _ = v16
                                _ = v17
                            }
                        }
                    case "foo bar":
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 uint64 = 486
                            v20 := GenS1{F0: 912}
                            v21 := GenS1{F0: -1}
                            _ = fi18
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    case "":
                        switch int(v13) {
                        case 127:
                            v7 = v7
                        case 32768:
                            fmt.Println(v2, v3)
                        default:
                            v22 := GenS2{F0: (^(fi15))}
                            _ = v22
                        }
                    }
                case 2147483647:
                    var v23 uint64 = 2
                    _ = v23
                case 641:
                    fmt.Println(v3, v7, fi15)
                }
                _ = fi15
            }
        }
    }
    v24 := GenS1{F0: 32767}
    v25 := GenS1{F0: -32768}
    if false {
        if (-368 >= v13) {
            v2 -= (v2 | (v2 ^ -2147483648))
            switch -9223372036854775808 {
            case 217:
                var v26 float32 = 1.5
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 4; fi27++ {
                for fi28 := 0; fi28 < 3; fi28++ {
                    fmt.Println(fi28, v13)
                    var v29 bool = (true || false)
                    _ = fi28
                    _ = v29
                }
                _ = len(v4)
                _ = fi27
            }
            var v30 float64 = 1e10
            _ = v30
        }
        fmt.Println(v2, v3)
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := GenS1{F0: -656}
            fmt.Println(v13, v13)
            _ = fi31
            _ = v32
        }
    }
    var v33 int32 = -184
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v7
}
