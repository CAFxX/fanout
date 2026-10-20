package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int = 234
    v2 := []string{"hello"}
    v3 := make(map[int]string)
    var v4 int64 = 1
    var v5 int16 = -1
    var v6 int = v1
    v5 = (^(v5))
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 2}
        v5 += -32768
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 32767}
    v10 := GenS1{F0: 1}
    switch ("" + "a") {
    case "\n\t":
        v1 = (v6 >> ((735) & 63))
    case "":
        switch (^(v1)) {
        case -2147483648:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case -733:
            var v11 int16 = 1000
            _ = v11
        case 0:
            switch "hello" {
            case "\n\t":
                var v12 string = ("\n\t" + "")
                _ = v12
            case "foo bar":
                fmt.Println(v6)
            case "a":
                if true {
                    v3[(v6 ^ 32767)] = ("\n\t" + "foo bar")
                    var v13 int8 = -1
                    v1 += 128
                    _ = v13
                } else {
                    delete(v3, v1)
                    v14 := GenS1{F0: 18}
                    _ = v14
                }
            }
        default:
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 int64 = (-(v4))
                v3[fi15] = ""
                var v17 int64 = v4
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    case "hello":
        v3[v1] = ""
    default:
        var v18 int64 = (v4 + 42)
        _ = v18
    }
    if ((v5 < (-(v5))) || (15 >= 0)) {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v4)
        case "a":
            _ = len(v2)
        case "foo bar":
            if (!false) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := GenS2{F0: (1.5 <= 0.0)}
                    for fi21 := 0; fi21 < 8; fi21++ {
                        var v22 uint16 = 42
                        _ = fi21
                        _ = v22
                    }
                    var v23 int32 = 0
                    _ = fi19
                    _ = v20
                    _ = v23
                }
                var v24 uint64 = 42
                _ = v24
            } else {
                v1 += v6
                v25 := GenS2{F0: true}
                _ = v25
            }
        }
    } else {
        fmt.Println(v6)
    }
    v5 = v5
    var v26 int64 = v4
    var v27 int16 = v5
    var v28 string = "foo bar"
    _ = len(v2)
    v29 := GenS1{F0: 65535}
    switch (v28 + ("foo bar" + "hello")) {
    case "hello":
        v26 = v26
    case "a":
        if (v28 == "\n\t") {
            var v30 uint64 = 0
            v5 += v27
            if (v1 <= v1) {
                v30 = v30
                fmt.Println(v4, v1, v1)
                for fi31 := 0; fi31 < 19; fi31++ {
                    var v32 uint32 = 4294967295
                    v27 += (^(v27))
                    _ = fi31
                    _ = v32
                }
            } else {
                v5 += v5
                v28 = ("foo bar" + v28)
                fmt.Println(v27, v6)
            }
            _ = v30
        }
    default:
        v2 = append(v2, ("" + "foo bar"))
    }
    v27 = (-(v27))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
