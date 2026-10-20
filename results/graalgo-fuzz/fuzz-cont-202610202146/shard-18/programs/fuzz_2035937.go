package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{789, 831, 907}
    var v2 int32 = -793
    var v3 float32 = (-2.25 * 1.5)
    v4 := make(map[int]string)
    v3 = v3
    v2 += 2147483647
    v5 := GenS1{F0: -9223372036854775808}
    if len(v1) > 0 {
        v1[0] = 2
    }
    v2 = v2
    if (42 == 406) {
        var v6 int64 = 0
        v2 += 42
        if false {
            if (!(!true)) {
                v7 := GenS1{F0: v6}
                v3 = 0.1
                _ = v7
            }
        } else {
            var v8 uint32 = 661
            v9 := v4[1]
            _ = v8
            _ = v9
        }
        _ = v6
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: 1}
        for fi12 := 0; fi12 < 11; fi12++ {
            switch fi12 {
            case 128:
                switch -997 {
                case 930:
                    v2 += v2
                case 611:
                    v13 := GenS1{F0: 9223372036854775807}
                    _ = v13
                case -2147483648:
                    switch "a" {
                    case "a":
                        fmt.Println(v3)
                    default:
                        var v14 uint = 39
                        _ = v14
                    }
                }
            case 4294967295:
                fmt.Println(v3, fi12)
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        if (833 == 0) {
            var v16 uint64 = 2
            if true {
                v17 := v4[(^(fi15))]
                v18 := v4[(-(fi15))]
                v19 := GenS1{F0: -1}
                _ = v17
                _ = v18
                _ = v19
            } else {
                var v20 uint32 = 2
                switch (("hello" + ("a" + ("" + ("a" + ((("\n\t" + "a") + "a") + ""))))) + "") {
                case "a":
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v20 = v20
                        _ = fi21
                    }
                case "hello":
                    var v22 int = fi15
                    _ = v22
                }
                _ = v20
            }
            _ = v16
        }
        _ = fi15
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 uint8 = 25
        v24 = (v24 >> ((1) & 63))
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
