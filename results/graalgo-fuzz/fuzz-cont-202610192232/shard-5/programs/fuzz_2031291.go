package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 string = ""
    v2 := []string{v1}
    var v3 uint32 = 4294967295
    var v4 uint8 = 2
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v8 := GenS1{F0: (v4 >> ((65) & 63))}
            var v9 float32 = 3.14159
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 int16 = -89
    if true {
        v1 = v1
    }
    var v11 uint64 = 560
    switch ((v1 + (("" + (v1 + (v1 + "a"))) + v1)) + "hello") {
    case "":
        for fi12 := 0; fi12 < 9; fi12++ {
            var v13 int16 = (v10 & (v10 / ((int16(fi12)) | 1)))
            v4 = (v4 >> ((794) & 63))
            _ = fi12
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int8 = 100
                fmt.Println(fi14)
                switch fi14 {
                case 1:
                    var v17 int = (fi15 << ((4294967295) & 63))
                    _ = v17
                }
                v5[uint32(v4)] = "a"
                v3 = 42
                _ = fi15
                _ = v16
            }
            fmt.Println(v4, v3)
            _ = fi14
        }
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS1{F0: (v4 + v4)}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v5[v3] = v1
    v4 = (v4 & 154)
    v21 := GenS1{F0: v4}
    switch 221 {
    case 32768:
        switch 2147483647 {
        case -885:
            for fi22 := 0; fi22 < 14; fi22++ {
                v3 = (^(v3))
                _ = fi22
            }
        case 328:
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(v1)
                v24 := GenS2{F0: 2}
                _ = fi23
                _ = v24
            }
        }
    case -2147483648:
        if (1 <= 609) {
            for fi25 := 0; fi25 < 13; fi25++ {
                v11 = 1
                _ = fi25
            }
        }
    case -41:
        switch v1 {
        case "":
            switch (v1 + (v1 + ((v1 + "foo bar") + v1))) {
            case "":
                var v26 bool = (539 < 4294967295)
                _ = v26
            }
        case "hello":
            fmt.Println(v10, v1)
        default:
            fmt.Println(v4, v10, v11)
        }
    }
    v10 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
