package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int16 = -1
    var v4 uint16 = 32767
    var v5 string = "a"
    var v6 uint32 = 440
    v7 := GenS1{F0: (v3 - v3)}
    v8 := GenS1{F0: (v3 + v3)}
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 bool = ((-9223372036854775808 > 42) && false)
            if false {
                if ((!(v11 || v11)) && (!v11)) {
                    var v12 bool = (v11 || v11)
                    v13 := GenS1{F0: v3}
                    _ = v12
                    _ = v13
                }
            } else {
                var v14 int16 = (v3 ^ -385)
                v15 := GenS1{F0: v3}
                var v16 int16 = v14
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 uint = 646
            _ = fi10
            _ = v11
            _ = v17
        }
        _ = fi9
    }
    if (-9223372036854775808 <= 554) {
        if (v6 < 2147483648) {
            if false {
                for fi18 := 0; fi18 < 5; fi18++ {
                    if (!(false || (!(0 < 1)))) {
                        v19 := GenS1{F0: -82}
                        _ = v19
                    }
                    _ = fi18
                }
            } else {
                var v20 int32 = -139
                var v21 uint16 = 255
                _ = v20
                _ = v21
            }
        }
        v3 += (v3 >> ((855) & 63))
        if ((v5 == (((v5 + v5) + (v5 + v5)) + ("\n\t" + "foo bar"))) || false) {
            var v22 uint64 = 42
            _ = v22
        } else {
            switch "\n\t" {
            case "hello":
                var v23 uint16 = v4
                _ = v23
            case "foo bar":
                v3 += 1000
            case "":
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 uint8 = 154
                    _ = fi24
                    _ = v25
                }
            default:
                var v26 bool = (uint(v6) != 619)
                _ = v26
            }
        }
    } else {
        switch 357 {
        case 9223372036854775807:
            if true {
                fmt.Println(v6, v5)
                v5 = "\n\t"
                fmt.Println(v6, v4, v6)
            }
        case -154:
            var v27 uint8 = 49
            _ = v27
        case 32768:
            var v28 int32 = -731
            _ = v28
        }
        v3 = (v3 % ((v3) | 1))
    }
    if (0 == 812) {
        v4 = 65535
        v2[42] = (v6 >> ((42) & 63))
        fmt.Println(v5)
    } else {
        var v29 int64 = -481
        _ = v29
    }
    v3 = (v3 << ((uint(v4)) & 63))
    switch (v5 + v5) {
    case "hello":
        v30 := v1[127]
        _ = v30
    case "\n\t":
        fmt.Println(v6)
    }
    v31 := v2[32767]
    var v32 bool = true
    for fi33 := 0; fi33 < 3; fi33++ {
        v6 = (^(v6))
        v34 := GenS1{F0: 1}
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        if (74 != 985) {
            var v36 string = v5
            if false {
                v37 := GenS2{F0: 42}
                for fi38 := 0; fi38 < 10; fi38++ {
                    var v39 string = "hello"
                    _ = fi38
                    _ = v39
                }
                v40 := v2[0]
                _ = v37
                _ = v40
            }
            v41 := []int{fi35, int(v3), (fi35 - fi35), fi35}
            _ = v36
            _ = v41
        }
        for fi42 := 0; fi42 < 18; fi42++ {
            v43 := GenS3{F0: (v31 / ((v31) | 1))}
            var v44 uint8 = uint8(v4)
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
