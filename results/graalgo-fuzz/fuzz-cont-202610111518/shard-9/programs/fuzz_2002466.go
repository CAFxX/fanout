package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 32767
    var v3 int8 = 127
    v3 -= v3
    for fi4 := 0; fi4 < 9; fi4++ {
        switch 32767 {
        case 32767:
            v3 -= (v3 | v3)
        case -2147483648:
            if false {
                if false {
                    switch "foo bar" {
                    case "":
                        fmt.Println(v3, fi4)
                    case "hello":
                        for fi5 := 0; fi5 < 6; fi5++ {
                            v2 -= (v2 % (((^(v2))) | 1))
                            _ = fi5
                        }
                    }
                    if false {
                        switch (fi4 ^ fi4) {
                        case 4294967295:
                            if ((42 == 2147483648) || false) {
                                fmt.Println(fi4)
                            }
                        case -844:
                            var v6 int16 = (-(v2))
                            _ = v6
                        case 2147483647:
                            for fi7 := 0; fi7 < 9; fi7++ {
                                var v8 uint16 = 65535
                                _ = fi7
                                _ = v8
                            }
                        }
                        switch ("hello" + ("hello" + "\n\t")) {
                        case "a":
                            switch ((("" + "hello") + "hello") + "a") {
                            case "foo bar":
                                var v9 bool = (true && (0 != 42))
                                _ = v9
                            case "\n\t":
                                var v10 uint16 = 255
                                _ = v10
                            case "hello":
                                switch (("\n\t" + ("" + "a")) + "foo bar") {
                                case "":
                                    v11 := GenS1{F0: (1 > 255), F1: (v2 ^ (^(v2)))}
                                    _ = v11
                                }
                            }
                        case "foo bar":
                            v12 := []int{(fi4 * fi4), fi4}
                            _ = v12
                        case "\n\t":
                            v13 := GenS2{F0: (v2 * 32767), F1: 65535, F2: (fi4 - 0)}
                            _ = v13
                        }
                    } else {
                        v14 := []int{fi4, fi4, 888}
                        var v15 bool = (88 == 4294967295)
                        _ = v14
                        _ = v15
                    }
                }
            }
        }
        _ = fi4
    }
    var v16 bool = (4294967295 >= 785)
    v17 := GenS2{F0: (v2 >> ((0) & 63)), F1: 42, F2: 32767}
    var v18 int8 = (v3 >> ((0) & 63))
    for fi19 := 0; fi19 < 10; fi19++ {
        v16 = (-1 <= 2147483647)
        _ = fi19
    }
    v18 = v3
    v3 = (v3 | v3)
    v20 := GenS3{F0: 154, F1: (v18 + v3)}
    fmt.Println(v3, v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
