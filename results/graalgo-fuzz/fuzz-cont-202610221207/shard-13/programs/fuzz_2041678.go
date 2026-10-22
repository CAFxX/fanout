package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 142
    var v2 uint8 = 1
    v3 := make(map[string]int)
    var v4 int16 = 42
    v5 := []string{"\n\t", (("\n\t" + "foo bar") + "\n\t"), ""}
    var v6 uint = v1
    for fi7 := 0; fi7 < 3; fi7++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v6 = (^(v6))
        var v8 uint = v1
        _ = fi7
        _ = v8
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            v6 = v6
            _ = fi9
        }
    } else {
        v1 = 2
        var v10 int = -997
        _ = v10
    }
    var v11 int8 = 42
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 2}
        fmt.Println(v2, v6)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            v4 += -1
            fmt.Println(v11)
            if (-1 >= (-(v11))) {
                if true {
                    v16 := GenS1{F0: 0}
                    v5 = append(v5, "")
                    _ = v16
                }
                v1 += (^(v1))
            } else {
                v17 := GenS1{F0: 255}
                _ = v17
            }
            _ = fi15
        }
        _ = fi14
    }
    switch "\n\t" {
    case "foo bar":
        if (-9223372036854775808 >= -459) {
            var v18 int32 = -902
            _ = v18
        }
    }
    if (!false) {
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 uint32 = 845
            _ = fi19
            _ = v20
        }
        if false {
            for fi21 := 0; fi21 < 14; fi21++ {
                switch "hello" {
                case "a":
                    var v22 int64 = -9223372036854775808
                    _ = v22
                case "":
                    var v23 int8 = -3
                    _ = v23
                case "hello":
                    fmt.Println(v1, v2, v6)
                }
                _ = fi21
            }
            v3[("hello" + ("hello" + (("a" + "hello") + "hello")))] = -958
            v6 = v6
        } else {
            var v24 bool = true
            fmt.Println(v24, v6)
            var v25 string = "foo bar"
            _ = v24
            _ = v25
        }
        v4 = v4
    } else {
        v26 := GenS2{F0: 2}
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 string = ""
            delete(v3, "")
            switch (fi27 * fi27) {
            case 127:
                v2 += v2
            case -524:
                for fi29 := 0; fi29 < 10; fi29++ {
                    v11 = (v11 & int8(v4))
                    v4 -= int16(fi27)
                    var v30 float32 = 3.14159
                    _ = fi29
                    _ = v30
                }
            case 338:
                v1 += (v1 >> (((v6 * v1)) & 63))
            }
            _ = fi27
            _ = v28
        }
        if (!(773 > 969)) {
            var v31 bool = true
            if true {
                var v32 string = "a"
                for fi33 := 0; fi33 < 8; fi33++ {
                    var v34 int8 = -128
                    v4 = 1
                    fmt.Println(v34, v34, v32)
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
            _ = v31
        }
        _ = v26
    }
    var v35 uint = 1
    var v36 uint32 = 2
    v37 := GenS1{F0: 569}
    switch "hello" {
    case "\n\t":
        var v38 int64 = 9223372036854775807
        _ = v38
    case "":
        var v39 uint16 = 0
        _ = v39
    }
    var v40 int = 60
    _ = len(v3)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
